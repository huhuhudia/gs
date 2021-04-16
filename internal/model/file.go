package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type File struct {
	FullPath    string `json:"full_path"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Dir         string `json:"dir"`
	Content     []byte
	ContentStr  *string
	FindTargets []*Target
}

type Target struct {
	LineNum     int    `json:"line_num"`
	LineContent string `json:"line_content"`
}

func (f *File) DebugInfo() string {
	return f.FullPath
}

func (f *File) GetContent() {
	content, err := ioutil.ReadFile(f.FullPath)
	if err != nil {
		logrus.Errorln(f.DebugInfo(), "read file error ")
		return
	}
	f.Content = content
}

func (f *File) FindTarget(target string) {
	if len(*f.ContentStr) == 0 {
		tmp := string(f.Content)
		f.ContentStr = &tmp
	}
}

func (f *File) PrettyInfo() {
	fmt.Println("===================================================")
	fmt.Println(f.FullPath + ":1")
	for _, target := range f.FindTargets {
		fmt.Printf("%s:%d %s \n", f.FullPath, target.LineNum, target.LineContent)
	}
	fmt.Println("===================================================")
}
