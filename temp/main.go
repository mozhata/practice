package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
)

func main() {
	fakeRequest()
}

func fakeRequest() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./log_dir")
	flag.Parse()
	glog.Infoln("begin fake..")
	// test_uid := "45529"
	host := "0.0.0.0"
	port := "12231"
	dirver_trail_json := `[{"lat":"33.33066","lng":"121.284148","t":1472338663}]`
	driver_trail_form := url.Values{
		"session_id": []string{"test_uid"},
		"json":       []string{dirver_trail_json},
		"city":       []string{"上海"},
	}

	ticker := time.NewTicker(time.Second * 5)
	for t := range ticker.C {
		resp, err := http.PostForm(fmt.Sprintf("http://%s:%s/driver/trail", host, port), driver_trail_form)
		if err != nil {
			glog.Errorf("at %s, PostForm-err: %s\n", t, err)
		}
		defer resp.Body.Close()
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Errorf("at %s, ioutil.ReadAll-err: %s\n", t, err)
		}
		glog.Infof("at %s, fake post, fadback is %s", t, string(resp_body))
	}
}
