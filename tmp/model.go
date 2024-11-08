package tmp

func ModelTmp() string {
	return `package {{.Package}}

import (
	"github.com/sonhineboy/gsadmin/service/global"
)

type {{.Name}} struct {
	global.GAD_MODEL
	{{range .Fields}}{{. | transField}}
	{{end}}
}

func (m *{{.Name}}) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "{{.Table}}")
}
`

}
