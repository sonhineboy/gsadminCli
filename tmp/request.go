package tmp

func RequestTmp() string {
	return `package {{.Package}}
type {{.Name}}Request struct {
}
`
}
