package image

import (
	"bytes"
	"image"
	_ "image/png"
	"io"
	"os"

	"practice/go/util"
)

func TryDecode(file *os.File) {
	var buf bytes.Buffer
	m, s, err := image.Decode(io.TeeReader(file, &buf))
	util.Debug("err is nil: %t, &err is nil: %t", err == nil, &err == nil)
	if err != nil {
		util.CheckErr(err)
	}
	reader := bytes.NewReader(buf.Bytes())

	_, _ = m, reader
	util.Debug("s: %s", s)
}
