package basecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/mozhata/merr"
)

type Controller struct {
	beego.Controller
}

type Resp struct {
	merr.MErr
	Data interface{} `json:"data"`
}

func (c *Controller) HandleErr(err error) {
	e := merr.WrapErr(err)
	glog.Errorf("err: %s\nraw err msg: %s\ncall stack: %s", e.Error(), e.RawErr().Error(), e.CallStack())
	body := Resp{
		MErr: *e,
	}
	c.Response(e.StatusCode, body)
}

func (c *Controller) Success(data interface{}) {
	body := Resp{
		MErr: merr.MErr{
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
