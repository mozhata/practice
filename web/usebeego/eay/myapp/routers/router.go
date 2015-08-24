package routers

import (
	"practice/web/usebeego/eay/myapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
