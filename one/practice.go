package main

import (
	// "flag"
	"fmt"
	"time"
	// "github.com/PuerkitoBio/goquery"
	// "bitbucket.org/applysquare/applysquare-go/cmd/catcher/universial"
	"github.com/zykzhang/goSnippet"
	// "os"
	// "path"
)

// var CurPath, _ = os.Getwd()

func main() {
	// flag.Parse()
	// fmt.Println("CurPath: ", CurPath)
	// fmt.Println("hello")
	fmt.Println(goSnippet.CurPath())
	fmt.Println(time.Now())
	// goSnippet.PrintFileBase()
	// goSnippet.Hello()
	// fmt.Println(goSnippet.GetUserHomeDir())
	// goSnippet.TestGlog()
	// goSnippet.TestReadCsv()

	// goSnippet.TryReadCsvByLine()
	// fmt.Println("........")
	// handler := goSnippet.GetOneline()
	// for i := 0; i < 9; i++ {
	// 	fmt.Println(handler())
	// }
	// fmt.Println("........")
	// handler = goSnippet.GetOneline()
	// for {
	// 	if str := handler(); len(str) <= 0 {
	// 		break
	// 	}
	// 	fmt.Println(handler())
	// }
	// goSnippet.WeiduOs()
	// goSnippet.TryHttp()
	// goSnippet.TryChan()
	// goSnippet.TryRegex()
	// goSnippet.TryReplace()
	// goSnippet.TryMap()
	goSnippet.TesLowerFunc()
	// goSnippet.Travel()
	// goSnippet.Slice()
	// var (
	// 	v1 int     = 2
	// 	v2 int64   = 424
	// 	v3 string  = "hello"
	// 	v4 float32 = 23.0
	// )
	// goSnippet.MyPrintf(v1, v2, v3, v4)
	// goSnippet.Closure()
	// goSnippet.TryChannel()
	goSnippet.Reflec()

}
