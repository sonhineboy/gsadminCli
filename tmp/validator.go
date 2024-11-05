package tmp

func ValidatorTmp() string {
	return `package {{.Package}}

import (
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
)

type {{.Name}}Validator struct {
	ginValidator.BaseValidator
}

func (d *{{.Name}}Validator) TagName() string {
	return "{{.Name | }}"
}

func (d *{{.Name}}Validator) Messages() string {
    //This is error message
	return ""
}

func (d *{{.Name}}Validator) Validator(fl validator.FieldLevel) bool {
	//To Do .....
	return true
}
`
}
