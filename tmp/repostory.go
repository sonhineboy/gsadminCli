package tmp

func RepositoryTmp() string {

	return `package {{.Package}}

type {{.Name}}Repository struct {
	BaseRepository
}
// New{{.Name}}Repository 实例化
func New{{.Name}}Repository() *{{.Name}}Repository {
	return &{{.Name}}Repository{}
}
`
}
