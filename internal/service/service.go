package service

import (
	"github.com/huhuhudia/gs/internal/model"
	"github.com/huhuhudia/gs/internal/walk"
	"sort"
	"sync"
)

type FileAccessor struct {
	rootPath string
	files []*model.File
	limits chan struct{}
}
func (a *FileAccessor) Init(goNum int, rootPath string){
	a.limits = make(chan struct{}, goNum)
	a.rootPath =  rootPath
}

func (a *FileAccessor) WalkDir(){
	a.files = walk.GetAll(a.rootPath)
}
func (a *FileAccessor) GetTargets(target string){
	model.Init(target)
	wg := &sync.WaitGroup{}
	for i:=0; i<len(a.files); i++{
		a.limits<- struct{}{}
		wg.Add(1)
		go func(idx int){
			a.files[idx].GetContent()
			a.files[idx].FindTargets = model.FindIndex(a.files[idx].Content)
			wg.Done()
			<-a.limits
		}(i)
	}
	wg.Wait()
	sort.SliceStable(a.files, func(i, j int) bool {
		return a.files[i].FullPath < a.files[j].FullPath
	})

}

func (a *FileAccessor)PrettyInfo(){
	for i:=0; i<len(a.files); i++{
		if len(a.files[i].FindTargets) != 0{
			a.files[i].PrettyInfo()
		}
	}
}
