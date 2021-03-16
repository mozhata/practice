package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"practice/go/file/image"
	"practice/go/util"
	"syscall"

	"github.com/mozhata/merr"
)

const (
	imagePath = "./testdata/cc.png"
	filePath  = "./testdata/password"
)

func main() {
	// trySeek()
	// tryDecodeImage()
	// countDirSize("/Users/mozhata/go/src/practice")
	// countDirSize("/Users/mozhata/go/src/practice/")
	// countDirSize("/Users/mozhata/go/src")
	// countDirSize("/Users/mozhata/go/src/")
	// tryAbs()
	tryChannelClose()
}

// https://www.jianshu.com/p/d24dfbb33781
func tryChannelClose() {
	type T int
	err := errors.Unwrap(fmt.Errorf("this is er wrap %s", errors.New("lll")))
	errors.Is(err, os.ErrExist)

	var IsClosed = func(ch <-chan T) bool {
		select {
		case <-ch:
			return true
		default:
		}

		return false
	}
	c := make(chan T)
	fmt.Println(IsClosed(c))
	close(c)
	fmt.Println(IsClosed(c))
	fmt.Println(IsClosed(c))

	_ = IsClosed
}

func tryAbs() {
	fmt.Println(os.Getwd())
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Abs("../testdata/cc.png"))
	fmt.Println(filepath.Abs("/testdata/cc.png"))
	fmt.Println(filepath.Abs("./../testdata/cc.png"))
	fmt.Println(filepath.Clean("/test/*"))
	dir := "/Users/mozhata/go/src/practice/go/file"
	dir = filepath.Dir(dir)
	fmt.Println(dir)
	dir = filepath.Dir(dir)
	fmt.Println(dir)
	dir = filepath.Dir(dir)
	fmt.Println(dir)
	fi, err := os.Stat("testdata/cc.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("fileName: %s\n", fi.Name())
	fi, err = os.Stat("testdata")
	if err != nil {
		panic(err)
	}
	fmt.Printf("fileName: %s\n", fi.Name())
	if err := os.MkdirAll("/tmp/gotest/", 0400); err != nil {
		panic(err)
	}
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	if err := os.MkdirAll("/tmp/gotest2/", 0400); err != nil {
		panic(err)
	}
}

func tryDecodeImage() {
	f, err := os.Open(imagePath)
	util.CheckErr(err)
	image.TryDecode(f)
}

func trySeek() {
	f, err := os.Open(imagePath)
	util.CheckErr(err)
	size, err := f.Seek(0, 2)
	util.CheckErr(err)
	util.Debug("size: %d", size)
}

func countDirSize(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	st, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s size is %d\n", path, st.Size())
}
