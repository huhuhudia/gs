package walk

import (
	"github.com/huhuhudia/gs/internal/model"
	"os"
	"path/filepath"
	"strings"
)

func GetAll(path string)(res []*model.File){

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".git")||
			strings.Contains(path, ".idea")||
			info.IsDir()||
			!strings.Contains(path, "."){
			return nil
		}
		res = append(res, &model.File{
			FullPath: path,
		})
		return nil
	})
	return res
}