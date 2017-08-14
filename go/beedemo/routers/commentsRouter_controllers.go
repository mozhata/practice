package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"] = append(beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"] = append(beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"] = append(beego.GlobalControllerRouter["practice/go/beedemo/controllers:UserController"],
		beego.ControllerComments{
			Method: "CheckExistence",
			Router: `/:uname/existance`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
