package basecontroller

import (
	"encoding/json"
	"net/http"
	"practice/go/beedemo/util"

	"github.com/golang/glog"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

type Resp struct {
	Data interface{} `json:"data"`
	util.BaseErr
}

func (c *Controller) HandleErr(err error) {
	e := util.WrapErr(err)
	glog.Errorf("err: %s\nstckTrace: %s", e.Error(), e.Stack())
	body := Resp{
		BaseErr: *e,
	}
	c.Response(e.StatusCode, body)
}

func (c *Controller) Success(data interface{}) {
	body := Resp{
		Data: data,
		BaseErr: util.BaseErr{
			StatusCode: http.StatusOK,
			Message:    "success",
		},
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
