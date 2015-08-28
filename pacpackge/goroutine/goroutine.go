package main

import (
	"fmt"
	"time"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// basic()
	// Timeout()
	classicalMod()
}
func basic() {
	chs := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		// why need this ?
		chs[i] = make(chan int)
		go func(ch chan int) {
			P("counting...")
			ch <- 1
		}(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
func Timeout() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		P("one")
	case <-timeout:
		P("timeout !")
	}
}
func classicalMod() {
	// 用一个goroutine来准备工作,用另一个goroutine来执行处理,
	// 主goroutine和一些通道来安排其他事情
	jobList := []string{"a", "b", "c", "d"}

	jobs := make(chan string)
	done := make(chan bool, len(jobList))

	// 准备工作
	go func() {
		for _, job := range jobList {
			jobs <- job
		}
		close(jobs)
	}()

	// 执行处理
	go func() {
		for job := range jobs {
			P("job: ", job)
			done <- true
		}
	}()

	// 主goroutine 安排其他事情,确保主goroutine等到所有工作完成后才退出
	for i := 0; i < len(jobList); i++ {
		<-done
	}

}
