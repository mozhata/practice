package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("appname: " + beego.AppConfig.String("appname") +
		"\nhttpport: " + beego.AppConfig.String("httpport") +
		"\nrunmode: " + beego.AppConfig.String("runmode"))

	c.Ctx.WriteString("\n\nappname: " + beego.AppName +
		"\nhttpport: " + strconv.Itoa(beego.HttpPort) +
		"\nrunmode: " + beego.RunMode)

	beego.Trace("trace test1..")
	beego.Info("test1...")

	beego.SetLevel(beego.LevelInformational)

	beego.Trace("trace test2..")
	beego.Info("test2...")
}
