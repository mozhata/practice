package main

import (
	"fmt"
	"path"

	"bitbucket.org/applysquare/applysquare-go/cmd/catcher/universial"
)

func main() {
	fmt.Println(universial.IsMail("http://www.baidu.com"))
	fmt.Println(universial.GetUserHomeDir())
	fmt.Println(path.Clean(universial.GetUserHomeDir() + "/test//b/as//dg"))
}
