package basecontroller

import (
	"encoding/json"
	"net/http"

	"practice/go/beedemo/util"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type Controller struct {
	beego.Controller
}

type Resp struct {
	util.BaseErr
	Data interface{} `json:"data"`
}

func (c *Controller) HandleErr(err error) {
	e := util.WrapErr(err)
	glog.Errorf("err: %s\norigin err msg: %s\ncall stack: %s", e.Error(), e.OriginErr(), e.CallStack())
	body := Resp{
		BaseErr: *e,
	}
	c.Response(e.StatusCode, body)
}

func (c *Controller) Success(data interface{}) {
	body := Resp{
		BaseErr: util.BaseErr{
			StatusCode: http.StatusOK,
			Message:    "success",
		},
		Data: data,
	}
	c.Response(http.StatusOK, body)
}

func (c *Controller) Response(statusCode int, data interface{}) {
	if data == nil {
		c.Ctx.ResponseWriter.WriteHeader(statusCode)
		return
	}
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.WriteHeader(statusCode)
	c.Ctx.ResponseWriter.Write(body)
}
