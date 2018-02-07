package main

import (
	"os"
	"practice/go/file/image"
	"practice/go/util"
)

const (
	imagePath = "./testdata/cc.png"
	filepath  = "./testdata/password"
)

func main() {
	// trySeek()
	tryDecodeImage()
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
