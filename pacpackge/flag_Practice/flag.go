package main

import (
	"flag"
	"fmt"
)

var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
func main() {
	// 定义一个string flag: -word, 其默认值为"foo"
	// flag.String函数返回值为指针类型
	wordPtr := flag.String("word", "foo", "intro: a string")
	numbPtr := flag.Int("num", 33, "an int")
	boolPtr := flag.Bool("fork", false, "an int")

	// 也可以使用变量
	// 定义一个flag -dada 存入svar中
	var svar string
	flag.StringVar(&svar, "dada", "dada", "a string var")

	// flag 定义完之后需要用flag.Parse()函数解析一下
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// 返回flag的个数
	fmt.Println("NFlag: ", flag.NFlag())
	// 去除flag之外的其他参数的个数
	fmt.Println("NArg: ", flag.NArg())

	if flag.NArg() > 0 {
		fmt.Println("args2:", flag.Args()[0])
	}

}

// flag 参数的位置不必按照顺序来,若不给flag传参数,则使用其默认值:
// zyk@zyk-pc:flag_Practice$ go run flag.go -word=off -dada=dadada~ -num=99 -fork=true asgasdgal sdaga
// word: off
// num: 99
// fork: true
// svar: dadada~
// tail: [asgasdgal sdaga]
