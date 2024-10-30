package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type FileOperator struct {
	AllPath, Dir, Name, Ext string
}

func (o *FileOperator) DirExist() bool {
	_, err := os.Stat(o.AllPath)
	return err == nil
}

func (o *FileOperator) CreateDir() error {
	return os.MkdirAll(o.Dir, os.ModePerm)
}

func (o *FileOperator) ParseDir() {
	o.setDir(path.Dir(o.AllPath))
}

func (o *FileOperator) ParseName() {
	_, name := path.Split(o.AllPath)
	o.setName(strings.ReplaceAll(name, o.Ext, ""))
}

func (o *FileOperator) ParseExtAndSet() {
	o.Ext = path.Ext(o.AllPath)
}

func (o *FileOperator) setDir(dir string) {
	o.Dir = dir
}

func (o *FileOperator) setName(name string) {
	o.Name = name
}

func (o *FileOperator) NameToTitle() string {
	return fmt.Sprint(strings.ToUpper(o.Name[:1]), o.Name[1:])
}
