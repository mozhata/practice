package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func main() {
	// encode
	hello := `iUxaTwANChoKEDAgQAlAAQUOMAA9AAAAAFaYvG1WmLxtADKLBZkAAAJmAAACNKu8
wBMAFzJhMzhhNGE5MzE2YzQ5ZTVhODMzNTE3YzQ1ZDMxMDcwIDg2MTM5ODVldAMA
E2I4Zjc1N2FlNjQzOWU4NzliYjJhIHwgUm9tYW4gSGl0bWFuCj0gHgAAAAcKL1Rk
NldGb0FBQVRtMXJSR0FnQWhBUndBQAAAAHRRejFqTTRBRnlBUlZkQUNrSUJFU2l3
VU5uMHF1VVVBRElIZ3kyCjVsU1RROFM3R1ZHcU1TeTBxNUlqbjR1cmFJS05zcjRW
Q2J1WmRhOXllTElEUWZYTkc2UUFpWW1YWmZXQmx5SjQKa2NJTGFHcnAyWkc0d1NY
dzUzaEJyOHNQd1RRb1B1QjNsTGtySEdWV2N5OFZxMklRUGlnaTU0d3dZNGpJVFpi
dgpudmwxOFNJYWtvOWsyZUhGd3h6SXY0b1BJWG1zeXRiQmRWZVVhTStFTzQwbm1J
UkthTk9SNzFNZUROZkxVQmFkCmdHa2RnM2t2alRPa0xwNkJiSWMzOS9icm1yTzBH
Z1ZhdUQ1WHRHVU5CdHJ2aTBlemRSSlJGcldxT1N1SkdKd2UKTWRsNVJrMTdidFc3
ekFUUmI5cXRYZTlDTW1JQUUvM2I0Y1d1WW5uKy9MdFFNYW1yU3BOWGk2VnpGMThZ
eEpMcwpycHZScTlRcTdyWk9ZZk10YXJxZWdRUG5TYVwwTDMAHHRqeEdhc0JiaThz
QUFiRUM4d0lBQUVvcVZQK3h4R2Y3CkFnQUFBQUFFV1ZvPQoRAAAAAAAA`
	debyte := base64Encode([]byte(hello))

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}
