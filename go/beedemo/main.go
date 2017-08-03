package main

import (
	"flag"
	_ "practice/go/beedemo/routers"

	"github.com/astaxie/beego"
)

func init() {
	flag.Parse()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func main() {
	beego.Run()
}
