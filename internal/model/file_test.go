package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFile_GetContent(t *testing.T) {
	f := File{
		FullPath:"/Users/xuwentao/work/gs/go.mod1",
	}
	f.GetContent()
	fmt.Println(string(f.Content))
}

func TestFile_FindTarget(t *testing.T) {
	Init("logrus")
	f := File{
		FullPath:"/Users/xuwentao/work/gs/internal/model/file.go",
	}
	f.GetContent()
	res := FindIndex(f.Content)
	for _, tmp := range res{
		tmp, _ := json.Marshal(tmp)
		fmt.Println(string(tmp))
		f.FindTargets = res
	}
	f.PrettyInfo()
}