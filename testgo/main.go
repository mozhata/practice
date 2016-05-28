package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	testLoop()
}

const (
	LOOPNUM = 10000
)

func testLoop() {
	glog.Infof("begin, LoopNum: %d", LOOPNUM)
	beginTime := time.Now()
	{
		subBeginTime := time.Now()
		for i := 0; i < LOOPNUM; i++ {
			sum(1000000)
		}
		glog.Infof("a: spent time : %v", time.Now().Sub(subBeginTime))
	}
	{
		subBeginTime := time.Now()
		for i := 0; i < LOOPNUM*3; i++ {
			sum(1000000)
		}
		glog.Infof("b: spent time : %v", time.Now().Sub(subBeginTime))
	}
	{
		subBeginTime := time.Now()
		for i := 0; i < LOOPNUM*6; i++ {
			sum(1000000)
		}
		glog.Infof("c: spent time : %v", time.Now().Sub(subBeginTime))
	}
	{
		subBeginTime := time.Now()
		for i := 0; i < LOOPNUM*12; i++ {
			sum(1000000)
		}
		glog.V(2).Infof("d: spent time : %v", time.Now().Sub(subBeginTime))
	}

	glog.Infof("finished, spent time %v", time.Now().Sub(beginTime))
}
func sum(ceil int) {
	var sum int
	// glog.Infof("begin sum, ceil: %d", ceil)
	for i := 0; i < ceil; i++ {
		sum = sum + i
	}
	// glog.Infof("finish, sum: %d", sum)
}
