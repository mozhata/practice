package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// basic()
	// Timeout()
	// classicalMod()
	// classicalMod2()
	// trySelect()
	tryGMP()
}

/*
设置max=3, 开4个go runtime
使用ps 查看有几个进程
*/
func tryGMP() {
	runtime.GOMAXPROCS(3)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 1000; i < 2000; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 2000; i < 3000; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 3000; i < 4000; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(1000 * time.Second)
}

func trySelect() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("sed %d\n", i)
		}
		close(quit)
		fmt.Println("sed quit")
	}()
	for {
		select {
		case v := <-c:
			fmt.Printf("got %d, sleeping...\n", v)
			time.Sleep(time.Second * 2)
			fmt.Printf("sleep for %d done\n", v)
		case <-quit:
			fmt.Println("got exit sig, quit")
			return
		}
	}
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

	// for _, ch := range chs {
	// 	<-ch
	// }
	P("finish")
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

func classicalMod2() {
	// suffiexes, files := handleCommand()
	suffiexes := []string{"a", "ab", "s"}
	files := []string{"a.ab", "ba", "a.s", "b.s", "c.ab"}

	P(suffiexes, files)

	sink(filterSuffixes(suffiexes, source(files)))
}
func source(files []string) <-chan string {
	P("begin source..")
	out := make(chan string, 1000)

	go func() {
		for _, fileName := range files {
			out <- fileName
			P("transfered file: ", fileName)
		}
		close(out)
	}()
	return out
}
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	P("begin filterSuffixes...")
	out := make(chan string, 1000)

	go func() {
		for fileName := range in {
			if len(suffixes) == 0 {
				out <- fileName
				continue
			}

			ext := strings.ToLower(filepath.Ext(fileName))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- fileName
					P("deal with filename: ", fileName)
					break
				}
			}
		}
		close(out)
	}()
	return out
}
func sink(in <-chan string) {
	for fileName := range in {
		P(fileName)
	}
}
