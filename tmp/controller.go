package tmp

func ControllerTmp() string {
	return `package genExample
import (
	"github.com/gin-gonic/gin"
)

type {{.Name}}Controller struct{}

func (controller *{{.Name}}Controller) Index(ctx *gin.Context) {
}

func (controller *{{.Name}}Controller) Save(ctx *gin.Context) {

}

func (controller *{{.Name}}Controller) Edit(ctx *gin.Context) {
}

func (controller *{{.Name}}Controller) Delete(ctx *gin.Context) {

}

func (controller *{{.Name}}Controller) Get(ctx *gin.Context) {

}
`

}
