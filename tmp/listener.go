package tmp

func ListenerTmp() string {
	return `package {{.Package}}

import (
	"github.com/sonhineboy/gsadmin/service/pkg/event"
)

type {{.Name}}Listener struct {
}

func New{{.Name}}Listener() *{{.Name}}Listener {
	return &{{.Name}}Listener{}
}

func (l *{{.Name}}Listener) Process(e event.Event) {
}

`
}
