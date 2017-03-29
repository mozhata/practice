package main

import (
	"flag"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

// var baseURL = http://m.luoqiu.com/wapbook-34752_8_1/

func setFlag() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
}

func main() {
	setFlag()
	for i := 7; i < 8; i++ {
		// for i := 7; i < 70; i++ {
		url := fmt.Sprintf("http://m.luoqiu.com/wapbook-34752_%d_1/", i)
		glog.Infof("querying %s", url)
		doc, err := goquery.NewDocument(url)
		glog.Infoln("doc: ", doc.Find("html").Text())
		if err != nil {
			panic(fmt.Sprintf("failed to request url, index: %d, err: %s", err))
		}
		chaps := doc.Filter("body > .cover > .chapter")
		glog.Infoln("length: ", chaps.Length())
		chapLength := chaps.Length()
		for i := 0; i < chapLength; i++ {
			glog.Infoln("chap: ", chaps.Text())
		}
	}
}
