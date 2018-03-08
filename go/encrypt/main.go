package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/golang/glog"

	"practice/go/encrypt/account"
	"practice/go/encrypt/common"
	"practice/go/encrypt/db"
	"practice/go/encrypt/skeleton/route"
)

var (
	configPath = flag.String("conf", "conf/config.toml", "config file's path")
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	common.InitConfig(*configPath)
	db.InitMysql(&common.Config.MySQL)
}

func main() {

	regRoutes := account.NewRoute()

	handler := route.BuildHandler(
		regRoutes,
	)

	glog.Infof("start serving at %d", common.Config.Listen)
	if err := http.ListenAndServe(":"+strconv.Itoa(common.Config.Listen), handler); err != nil {
		panic(err)
	}

}
