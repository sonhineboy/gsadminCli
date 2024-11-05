package tmp

func EventTmp() string {
	return `package {{.Package}}

type {{.Name}}Event struct {
}

func New{{.Name}}Event() *{{.Name}}Event {
	return &{{.Name}}Event{
	}
}

func (e {{.Name}}Event) GetEventName() string {
	return "{{.Name}}Event"
}

`
}
