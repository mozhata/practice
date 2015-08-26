package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

// 几个现成的encoding:
// var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
// var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
// var StdEncoding = NewEncoding(encodeStd)
// var URLEncoding = NewEncoding(encodeURL)

var P func(...interface{}) (int, error) = fmt.Println
var input []byte = []byte(`
		123,
		abc,
		中文.`)

func main() {
	// tNewEncoder()
	// tDecodeString()
	// tEncodingToString()
	tNewEncoding()
}
func tNewEncoder() {
	// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
	// 使用enc方式加密
	// 加密后的字符串写入w
	// 必须关闭
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()
}
func tDecodeString() {
	str := "CgkJMTIzLAoJCWFiYywKCQnkuK3mlocu"
	// func (enc *Encoding) DecodeString(s string) ([]byte, error)
	// 解码返回[]byte类型
	data, _ := base64.StdEncoding.DecodeString(str)
	P("\n", "DecodeString: ", string(data))

	// string 和[]byte可以相互转换
	P([]byte("123"))
}
func tEncodingToString() {
	data := []byte("any + old & data zhong文.")
	str := base64.StdEncoding.EncodeToString(data)
	P(str)
}
func tNewEncoding() {
	base64Table := "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
	// P(len(base64Table))
	coder := base64.NewEncoding(base64Table)
	output := coder.EncodeToString(input)
	stoutput := base64.StdEncoding.EncodeToString(input)
	P(output)
	P(stoutput)

	input2, _ := coder.DecodeString(output)
	input2Std, _ := base64.StdEncoding.DecodeString(stoutput)
	P(string(input2))
	P(string(input2Std))

}
