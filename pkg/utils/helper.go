package utils

import (
	"os"
	"strings"
)

// GetPackage 路径的最后一个文件目录作为 包名
func GetPackage(dir string) string {
	dirs := strings.Split(dir, "/")
	return dirs[len(dirs)-1]
}

func CreateFilCallBack(allPath string, fn func(operatorFile *FileOperator, file *os.File) error) error {
	f := FileOperator{AllPath: allPath}
	f.ParseExtAndSet()
	f.ParseName()
	f.ParseDir()
	if !f.DirExist() {
		err := f.CreateDir()
		if err != nil {
			return err
		}
	}
	w, err := os.Create(f.AllPath)
	if err != nil {
		return err
	}
	defer func(w *os.File) {
		_ = w.Close()
	}(w)
	return fn(&f, w)
}
