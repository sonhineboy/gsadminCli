package utils

import "strings"

// GetPackage 路径的最后一个文件目录作为 包名
func GetPackage(dir string) string {
	dirs := strings.Split(dir, "/")
	return dirs[len(dirs)-1]
}
