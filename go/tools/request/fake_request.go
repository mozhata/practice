package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
)

const (
	BASE_RL            = "http://115.29.34.206:12236/driver/"
	TRAIL              = "trail"
	ORDER_TRIAL        = "orderTrail"
	TRIAL_PARAMS       = `?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`
	ORDER_TRIAL_PARAMS = `?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`
)

func main() {
	flag.Parse()
	FakeRequests()
}

func FakeRequests() {
	orderURI := fmt.Sprintf("%s%s%s", BASE_RL, TRAIL, TRIAL_PARAMS)
	orderTailURI := fmt.Sprintf("%s%s%s", BASE_RL, ORDER_TRIAL, ORDER_TRIAL_PARAMS)
	ticker := time.NewTicker(time.Second * 3)
	for t := range ticker.C {
		glog.Infof("fake request at %v", t)
		go func() {
			http.Get(orderURI)
			http.Post(orderURI, "", nil)
			http.Get(orderTailURI)
			http.Post(orderTailURI, "", nil)
		}()
	}
}
