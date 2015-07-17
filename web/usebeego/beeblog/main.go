package main

import (
	"practice/web/usebeego/beeblog/controllers"
	"practice/web/usebeego/beeblog/models"
	_ "practice/web/usebeego/beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", true, true)
	beego.Router("/", &controllers.MainController{})

	beego.Run()
}
