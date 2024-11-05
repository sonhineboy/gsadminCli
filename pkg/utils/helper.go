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

// CreateFileCallBack 创建文件回调
func CreateFileCallBack(allPath string, fn func(operatorFile *FileOperator, file *os.File) error) error {
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

// ToCamelCase 下划线风格转换大小驼峰风格
func ToCamelCase(s string, capitalizeFirst bool) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	result := strings.Join(words, "")

	if !capitalizeFirst && len(result) > 0 {
		// 如果需要小驼峰格式，将首字母改为小写
		result = strings.ToLower(result[:1]) + result[1:]
	}
	return result
}
