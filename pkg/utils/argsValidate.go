package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateFormat(str string) error {
	reg, _ := regexp.Compile("^\\./.*\\.go$")
	if !reg.MatchString(str) {
		return errors.New("格式必须以.go 结尾 以./开始")
	}
	return nil
}

func ValidateIsPackage(str string, pk string) error {
	if len(pk) == 0 && len(strings.Split(str, "/")) == 2 {
		return errors.New("根目录创建必须加上 --package")
	}
	return nil
}
