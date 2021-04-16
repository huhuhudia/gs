package main

import (
	"flag"
	"github.com/huhuhudia/gs/internal/service"
)

func main(){
	goNum := 1000
	rootPath := ""
	target := ""
	flag.StringVar(&rootPath, "path", "./", "search root path ")
	flag.IntVar(&goNum, "goNum", 1000, "concurrent handle file num")
	flag.StringVar(&target, "target", "", "target")

	flag.Parse()
	acessor := service.FileAccessor{}
	acessor.Init(goNum,rootPath)
	acessor.WalkDir()
	acessor.GetTargets(target)
	acessor.PrettyInfo()
}
