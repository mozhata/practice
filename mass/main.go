package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"github.com/mozhata/handy"
	"github.com/pborman/uuid"
)

var (
	P     func(...interface{}) (int, error)     = fmt.Println
	Pf                                          = fmt.Printf
	parse func(string, int, int) (int64, error) = strconv.ParseInt

	atoi         func(string) (int, error) = strconv.Atoi
	EmailPattern                           = `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})$`
)

type JSONTime time.Time

const (
	JSONTimeFormat = "2006-01-02T15:04:05Z"
)

var formated string = "2015-08-05T04:09:05Z"
var toConverted string = "2015-08-26T05:39:10Z"

type UMPlatform int

const (
	PlantformIos UMPlatform = iota
	PlantformAndriod
)

func main() {
	// boo, _ := regexp.MatchString(EmailPattern, "1.sg@q.com")
	// P("boo: ", boo)
	// P("format: ", time.Now().Format(JSONTimeFormat))
	// t, err := time.Parse(JSONTimeFormat, toConverted)
	// P("t: ", t, "\nerr : ", err)
	// t, err = time.Parse(JSONTimeFormat, formated)
	// P("t: ", t, "\nerr : ", err)
	// var s string
	// P(s == "")
	// P(formated[:len(formated)-1])
	// P(time.Now().Format("2006-01-02"))
	// P(rand.Int())
	// P(rand.Int())
	// P(PlantformIos)
	// P(PlantformAndriod)
	// source := rand.NewSource(8)
	// r := rand.New(source)
	// P(r.Int())
	// P(r.Int())
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// P(rand.Intn(5))
	// flo := float32(1.0 / 3)
	// P(flo)
	// P(time.Now())
	// P(time.Now().Local())
	// locationC, errC := time.LoadLocation("Asia/Shanghai")
	// locationU, errU := time.LoadLocation("UTC")
	// P(locationC, errC)
	// P(locationU, errU)
	// P(time.Now().In(locationC))
	// P(time.Now().In(locationU))
	// timeStamp := strconv.Itoa(int(time.Now().In(locationC).Unix()))
	// P(timeStamp)
	// v, e := atoi("a")
	// P(v, e)
	// v, e = atoi("")
	// P(v, e)
	// c := cron.New()
	// c.AddFunc("2 0 0 * * *", func() { fmt.Println("...") })
	// c.Start()
	// code := "abcd12"
	// P(code[:4])
	// password := fmt.Sprintf("%x", md5.Sum([]byte("cn.sduhaliluya")))[:4]
	// P(password)
	// var floaNan float64
	// var floatSmall float64 = 11.01
	// P(floaNan)
	// P(floatSmall)
	// P(math.IsNaN(floaNan), math.IsInf(floaNan, 0), math.IsInf(floaNan, 1), math.IsInf(floaNan, -1))
	// P(math.IsNaN(floatSmall), math.IsInf(floatSmall, 0), math.IsInf(floatSmall, 1), math.IsInf(floatSmall, -1))
	// P(&floaNan, &floatSmall)
	// var s map[string]interface{}
	// P(s)
	// P(len(s))
	// P(s == nil)
	// date := time.Now().Add(time.Hour * -24 * 7).Format("2006-01-02")
	// P(date)
	// var f float64
	// var ff float64 = 0.00001
	// P(f, ff)
	// P(ff > f)
	// P(ff > 0)
	// P(f > 0)
	// P(f == 0.0000000000000000000000)
	// var ss = struct {
	// 	a int
	// 	b string
	// }{}
	// P(ss)
	// P(&ss == nil)
	// l := field_of_study.KeyList
	// P(l)
	// P(len(l))

	// dic1 := map[string]string{"a": "aa", "as": "asas", "dd": "dddd"}
	// P(dic1, &dic1)
	// dic2 := dic1
	// dic2["a"] = "cc"
	// P(dic1, dic2)
	// P(&dic2)

	// var dic map[string]string
	// P(len(dic))
	// var s string
	// ll := strings.Split(s, ",")
	// P(ll)
	// P(len(ll))

	// reg := func() *regexp.Regexp {
	// 	wordList := []string{}
	// 	for i := range wordList {
	// 		wordList[i] = regexp.QuoteMeta(wordList[i])
	// 	}
	// 	return regexp.MustCompile(strings.Join(wordList, "|"))
	// }()
	// P(reg.MatchString("abc"))
	// P("reg is nil", reg == nil)
	// P(reg)
	// P(*reg)
	// P("string: ", reg.String(), reg.String() == "")
	// P(MarshalJSONOrDie(reg))
	// dest := []string{"abc", "dsa"}
	// fmt.Println(dest)
	// saveToSlice("source", dest)
	// fmt.Println(dest)
	// P(buldCountySlug(""))
	// P((1 == 2) ^ (3 == 4))
	// sl := make([]string, 0)
	// P(len(sl))
	// sle := make([]string, 7)
	// P(len(sle))
	// hello()

	// osPath()
	// urlParse()
	// loopMap()
	// sorttt()
	// crawl()
	// convertInterface()

	// runes()
	// logFlush()
	// build()
	// intLoop()
	// syncRuntime()
	// closure()
	// timeAdd()
	// tryCall()
	// getEnv()
	// tesMap()
	// tesSlice()
	// httpRequest()
	// tesErrorFmt()
	// tesArrayEqual()
	// testMapLen()
	// P(uuid.New())
	// WirdTest()
	// P(unsafe.Sizeof("s"))
	// P(unsafe.Sizeof(1))
	// P(unsafe.Sizeof(true))
	// testDefer()
	// tryBuffer()
	// tryPrintfV()
	// tryEmpty()
	// tryCrypto()
	// tiny()
	// tryPrintf()
	// tryLenStr()
	// tryUpper()
	// tryFileJoin()
	// makeMapLen()
	// graceGoruntine()
	// graceGoruntine()
	// tryCkecBankCard()
	// chanPrac()
	// chanPracBuffer()
	// formatFloat()
	// tryAddDate()
	// tryConvertPanicToError()
	// maxInt64()
	// tryIfElse()
	// lenStr()
	// tryUmarshal()
	// tryFeildFunc()
	// trySlice()
	// tryPem()
	// tryJson()
	// tryBreak()
	// tryRenameType()
	// tryDelv()
	// tryMethod()
	sendEmail()
}

// sen email
func sendEmail() {
	host := "smtp.163.com"
	userName := "zyk7676@163.com"
	PWD := "dx853556721"

	to := "mozhata@aliyun.com"
	targets := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	auth := smtp.PlainAuth("", userName, PWD, host)

	err := smtp.SendMail(host+":25", auth, userName, targets, msg)
	log.Printf("error of SendMail: %v\n", err)
}

func tryMethod() {
	m := M{make(map[string]interface{}), ""}
	fmt.Printf("origin m: %#v\n", m)
	m.modify("abc", "ABC")
	fmt.Printf("modified m: %#v\n", m)
	m.Modify("abc", "123")
	fmt.Printf("Modified m: %#v\n", m)
}

type M struct {
	Dic map[string]interface{}
	S   string
}

func (m M) modify(k, v string) {
	m.Dic[k] = v
	m.S = v
}
func (m *M) Modify(k, v string) {
	m.Dic[k] = v
	m.S = v
}

// 测试delve调试工具
func tryDelv() {
	dic := map[string]string{
		"abc": "ABC",
		"bcD": "BCD",
		"ASD": "ASD",
	}
	glog.Infof("dic is %#v\n", dic)
	for k, v := range dic {
		levelOne(k, v)
	}
}
func levelOne(k, v string) {
	fmt.Printf("level one, key: %s\tv: %s\n", k, v)
	levelTwo(k, v)
}
func levelTwo(k, v string) {
	fmt.Printf("level two, key=>value: %s => %s\n", k, v)
}

// 测试: 重命名之后的类型是否具有之前类型的方法
func tryRenameType() {
	type Req http.Request
	req := Req{}
	fmt.Println(req.Method) // 编译过

	// type T time.Time
	// t := T{}
	// fmt.Println(t.IsZero()) // 编译不过
}

func tryBreak() {
	for i := 0; i < 4; i++ {
		fmt.Println("the first level")
		for j := 0; j < 3; j++ {
			if i > 1 {
				fmt.Printf("i > 1, break (i: %d)\n", i)
				break
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}

func tryJson() {
	dic := map[string]string{
		"certificate": `-----BEGIN CERTIFICATE-----
MIIDBDCCAeygAwIBAgIRAMMu1b4BaHKrBZHlfnQYhMswDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAeFw0xNzAxMjUwNDA0NDRaFw0xODAxMjUwNDA0
NDRaMBIxEDAOBgNVBAoTB0FjbWUgQ28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCzjdtZ6KaY7RkcZYMr9awYySInCDkDy7wM/LRqQfoxd2JHs9ph5RMv
s3OPs033/qsAU+hS3EMCVH/D3hbwKdRZGdNv5i6xfq/H9cjbBybVcN8pGDc60lkD
6bFosKLTnwbg0XsEvGROR6+rxzA9BFXzEmHQW8gZDM4fPeg1c1MYzaw1IAKDxw3F
zKGCZdS+7SAgNVe7TkQNto/neIczLeB8uAnqcin/T4qBEE5TBZ2O6s7HqqUrBSv4
DWMDLr7ivqUXiTsJLECTuzF3oWJb8JP27LEsrRcE64EhR04vHFC48v4SjOXvtkhi
DJyBYRqc4WO1uWQXO0uXLV5ZESXeyWUHAgMBAAGjVTBTMA4GA1UdDwEB/wQEAwIF
oDATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMB4GA1UdEQQXMBWC
DXd3dy5iYWlkdS5jb22HBHp7ejcwDQYJKoZIhvcNAQELBQADggEBAHczTP2obQDy
VNbr3vmrV6KEdCQqMXjRHkMxLjR4PVUByVvQ1IBiv6l1p2+YFbKguPdcxX5e5ilC
5zqcMsX+HR7/gezKKy/rB1WXIxND4Kaf2iT8NZotT4qZn8TruCKgbEQp++Zkoyyj
9GGMjZHz0O0nrjbxyo9FCz3pRmEcav6t/ZwwViCgga1b2IJDCHioEPpmwN1fBg1i
ccQJOg2G3KM/Xzrmceukl81nCdYvlsNhjBHiPGJDIuclRS9r4H6TxhXgI8u2VSXZ
U7tKiH2rwD7PriqJsEae/dmWyc57I4kKA0n/nHaTbxzBLqYCRkVNkvkI70EZZtAG
EmAaKqLGqLM=
-----END CERTIFICATE-----`,
		"private_key": `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAs43bWeimmO0ZHGWDK/WsGMkiJwg5A8u8DPy0akH6MXdiR7Pa
YeUTL7Nzj7NN9/6rAFPoUtxDAlR/w94W8CnUWRnTb+YusX6vx/XI2wcm1XDfKRg3
OtJZA+mxaLCi058G4NF7BLxkTkevq8cwPQRV8xJh0FvIGQzOHz3oNXNTGM2sNSAC
g8cNxcyhgmXUvu0gIDVXu05EDbaP53iHMy3gfLgJ6nIp/0+KgRBOUwWdjurOx6ql
KwUr+A1jAy6+4r6lF4k7CSxAk7sxd6FiW/CT9uyxLK0XBOuBIUdOLxxQuPL+Eozl
77ZIYgycgWEanOFjtblkFztLly1eWREl3sllBwIDAQABAoIBAAz7zgfDVgCq7Gzs
Pj74hdhI+wo9gvnuSovZ3iXs3W2kR1wN+lWn3epOHKYgfhvXs3/TkXIkcUgC2jZM
ofX94FgNEQf/ahL1qTpZYIvOCZXXkpuDH8NGIJ+yz1xVoYCR+Epbe3zNLqemcWKH
hcpzvM9V3fLWIDMHhQQnUtpCm5iob5CZnHyWycQq1dCWH0RGkeSz2/C3Y6javmh/
H9QMWp622bFrFBRZB6/d7fivZlWwFIZG/q60QPdeNo1HJVM0YSvYeEPHIRJ9Tvdx
JNeUpGke6B5dZGOueCjMOwnnoaJl+8dQgHU7kP3MuHFH7v6VWIGmPFH84PHK0DS3
1lyJzEECgYEA6pKsQhnozMPaUoROrUhg+7qJBj+17w4wtubFFxQILUiJZiqbH+wR
MuJP0L1qaXhdLkGlaQLVqbWzswAb8fkrkTfYZzxTCuF+foG75ksDV7OagoUU4GeJ
IVyPUWQGOkKcg1vlPGviTcncu1EpcX3TG0UP3KRcALoTvL6Yp8AWlyECgYEAw/Sb
V+2pZ0GMpOpSlXrgumVBG7O0L5puEJN/geKG9CHKP7CsOo/cgq3zw51rSWvUAOAm
yJIfNm+OJSeNm+XBTF3K2+PJHrRSxL8dRZMyRfIBCBHgxnczK1rRmgmouPPK4e3m
7oM7hdZ8ynInw5Anqt1PXyNKTzP6I0a80s1/fycCgYEAoGWMzmJDEok1p0j7N3gP
gPHLMm+sKvusCdUAncg/0x8PfMHTct+L3xxq8VQkCFyacr/GqGicy0AI0XRYp3v2
84SZP/Y63mzUfxKc9lsCvRx9oZP4c324ggx1n3Ti4UGdHiFfPZKTmxdWDOvxh74+
9R2jO/9TPdf/GQfHAKXTjKECgYAGvhDmSqAslF3RgtqKmCrJXxiJanqFtMuauGiH
wJKiLdN2s46JiU+uE2wyy+TYJuSpjwzmk9iNF/ONJbpCpfortYJ4ZHMUImJCsMzC
CxmMvJH6hBr22T0ifcJ+iVyL2J+ffH8Yr91JcqLimGDz4q7quyiy/lTdOs2djx6K
JlX/uwKBgERjsJADxCLTiMDVxfOQqYgMHbYC45e3NuP33udefVMnFy8c1vl7CrZp
CCCFdqzqU3pk8fgYUGuGELPHzDkO77YcfX1GxCYRN4Tpbsyk96j9/5UQlNSj1Pec
4ACUAm63lE6mZBUMqL//QzVvU3fi0HVGRx/XcTxdM62JM0Lwmx9s
-----END RSA PRIVATE KEY-----`,
	}
	fmt.Println(handy.MarshalJSONOrDie(dic))
}

func tryPem() {
	validKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA8hmSDarB6pGiN/GHmgaegdhk9hpYYigxx75Cs8gNnePEITc3
36e0DQlaQEgWBKVyKovchtyBkmM2RSY14/B2qe36FsPTA5xLOtUXgJaja3doWgGd
lRhhVVQxSCjgKFoxKcW1Jhyg31ZracuJD5vribiKFgzMD7qAKnrCD0l33Szms5x3
HicSskSdank+K0aEp+GBnw+ydDyJNSELLAy01uMb1wSLjJEswmL52c07w4nk3alX
kp/rMOd7CZTg5XGEVVz+w+kmsKVDx+uTQdnuYvBTgLjcTRJJDwFVIdKnQl8df0Qo
os5LoRL7zKPD2jlmpkCjkfdzM+boIoymQk695QIDAQABAoIBAQDjbzAtFLBta3jo
c0d/ZB5+PsrCBAfp1nbFfbBt8k8Qd8tlpNkhjAtyf8nAP7gnj0GMRK3cJ8S1v+xY
08empdzkUIMz5a0TZD2sPEmC7IEDUT6/52qhiTDPF3PgpT3HyssKwggmkJh0l1v7
HFFbT4/bZniatFTzpbIHtiEeeVdYd6UtXMEm59rZktJPvMgXt4+gfyD8kKoYtl9c
2ZUHcDvVzBtalDFxHXYRTn3kIJec2EDN+NHlwl0a8DiqkOa+X05l1JzsxLdF6D1t
f68ugv4p6peFxsEGPIAXFG4Cel/5kMbNx5ZLhzZyNOHhTl7FdfA+5vUgLxrY+GuI
cvpwRAUBAoGBAP+WHHlTv3yqKoSYyDqpLFpINTV5WQQASyefMMKBfGQWWDcL2ETw
n7z8lv3YL9JE70OEODYTVwcUZxoZnRA3uGrgjEalojoQP2D+hUbDKJSI7yKMK9AN
bcl52gNuG0EMzGwhhpU44bia5XiffR6/ZGGUSag5lA5fled9ZBChWcolAoGBAPJ9
3zNkzVfplUN/1/V7IFwP/A9KgoQEqaIjir6G1zSMO/f/No3yqgzR36nqQtrCM814
QTnrV/byNmiAdWqvIvcDseaN1M9eJkMH0OJNE/PjQeOVoi4UJq+tN+WiVYVA8TVb
kYIer5RlAWHnFy9WIa7y6EQrTSPFmIG8uF7+rHjBAoGBAJi1Itwm+rFMpszE0FWk
QJjMreX+U/49AqpwxdxhK2ZNp3V0QUqnjnmHXdvbcG8sutxXQpKkqYnUWenRRfOr
pydMdOO9ERmqHYQhbpYiArwEuQSmRYMwktkfUfBcuDH8qKMuwM+lxc/b1KFmgYZ+
ikZ3KC83/8s6t0ExvTjmftR1AoGBAKmYt16uhZ+S0r5ez5/0+XDqDRR0vEuxJyMr
UIJotGMgIWEVsgYdTimhL5kbKp7tbGWsUZI8s00Xok38aiKvUvkIMIhbcheGdaQz
9zgPSIEu1IpjwQYROe1sbMfCfaWxAgKbhG1JIOvcqNbcVS7aQrylTyU+U7lNHZi/
cQOfgukBAoGAf/SfoW7Syb0aLCxnkQYOI5uflq6ZAvzUcKWx5I46mvL+oZMP5PcV
ANFCvSA+d00Ldn/95bQqe44CR99y/T1rlThDek/h1LABlb5KEWOH0pQRjt4lhDKC
dLSyu3PWui6MWbiF2ASMqrnen4k5NqQfdQ8XYt3Yu8vijCMXgtYinBQ=
-----END RSA PRIVATE KEY-----`
	validCert := `-----BEGIN CERTIFICATE-----
MIIDAzCCAeugAwIBAgIQRE4TsVnPasdLrdenqmCeojANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMB4XDTE3MDEyNDA5NDgzMloXDTE4MDEyNDA5NDgz
MlowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAPIZkg2qweqRojfxh5oGnoHYZPYaWGIoMce+QrPIDZ3jxCE3N9+ntA0J
WkBIFgSlciqL3IbcgZJjNkUmNePwdqnt+hbD0wOcSzrVF4CWo2t3aFoBnZUYYVVU
MUgo4ChaMSnFtSYcoN9Wa2nLiQ+b64m4ihYMzA+6gCp6wg9Jd90s5rOcdx4nErJE
nWp5PitGhKfhgZ8PsnQ8iTUhCywMtNbjG9cEi4yRLMJi+dnNO8OJ5N2pV5Kf6zDn
ewmU4OVxhFVc/sPpJrClQ8frk0HZ7mLwU4C43E0SSQ8BVSHSp0JfHX9EKKLOS6ES
+8yjw9o5ZqZAo5H3czPm6CKMpkJOveUCAwEAAaNVMFMwDgYDVR0PAQH/BAQDAgWg
MBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwHgYDVR0RBBcwFYIN
d3d3LmJhaWR1LmNvbYcEhYJ1cDANBgkqhkiG9w0BAQsFAAOCAQEAZLoTyz8MLScS
iM2swPPKRDRgnbSLPIZEN6C1sXWiIsqhc9Osmitr9mC2jP/f4XCQClD2ECmXJzNq
xUuwt5uhbOyLFum5V01eJVm/nlo5BAziGPmKDeaxgDelhrKgxsezVvk3i9D9+xua
AaZdbvyMRDc+UN/OMKgTzYrkElPgAkgv6nKf46Fl6bEQViCOSRq1+KWJL18LbFL6
2HQPDNEY2aEJ69r/oCyjMdaipO09g1Ql8dp1xWyIov3zA62tx08C3iQ2v3UQ7grJ
UNMautOYTbNJlCqGLd1GMCErcKUQjcGg3x6F+EuA70o5Y1mS+IF1wYo06I18ASc7
76lakwpbHg==
-----END CERTIFICATE-----`
	// 	invalidKey := `-----BEGIN RSA PRIVATE KEY-----
	// MIIEpQI
	// -----END RSA PRIVATE KEY-----`
	// 	invalidCert := "FV9uaat2qwLfouya3"
	// var data []byte
	// var err error
	key, rest := pem.Decode([]byte(validKey))
	fmt.Printf("valid key of key: %s and rest: %s\n\n", handy.MarshalJSONOrDie(key), rest)
	cert, rest := pem.Decode([]byte(validCert))
	fmt.Printf("valid cert of cert: %s and rest: %s\n\n", handy.MarshalJSONOrDie(cert), rest)
	certificate, err := tls.X509KeyPair([]byte(validCert), []byte(validKey))
	fmt.Printf("err: %v\tcert:\n%s", err, handy.MarshalJSONOrDie(certificate))
	// fmt.Printf("byte: %s\n", base64.StdEncoding.EncodeToString(cert.Bytes))
	// _, err = encoding.Decode(data, cert.Bytes)
	// fmt.Println("data: ", string(data), err)

	// p, rest = pem.Decode([]byte(invalidKey))
	// fmt.Printf("invalid * of p: %s and rest: %s\n\n", handy.MarshalJSONOrDie(p), rest)
	// p, rest = pem.Decode([]byte(invalidCert))
	// fmt.Printf("invalid * of p: %s and rest: %s\n\n", handy.MarshalJSONOrDie(p), rest)
	// p, rest = pem.Decode([]byte(validKey + "\n" + validCert))
	// fmt.Printf("valid of bind of p: %s and rest: %s\n\n", handy.MarshalJSONOrDie(p), rest)
}

func trySlice() {
	var sl []string
	fmt.Println(len(sl), sl == nil)

	var m map[string]string
	fmt.Println(len(m), m == nil)

	// panic: use of untyped nil
	// fmt.Println(len(nil))

	fmt.Printf("joined:%q\n", strings.Join(sl, ","))

	sl = make([]string, 3, 5)
	sl[3] = "3"
	fmt.Printf("%v\n", sl)
}

func tryFeildFunc() {
	url1 := "http://192.168.1.123:8989"
	url2 := "http://192.168.1.123"
	url3 := "http://192.168.1.123:"
	split := func(s rune) bool {
		return s == '/' || s == ':'
	}
	fmt.Printf("parts of url1: %#v\n", strings.FieldsFunc(url1, split))
	fmt.Printf("parts of url2: %#v\n", strings.FieldsFunc(url2, split))
	fmt.Printf("parts of url3: %#v\n", strings.FieldsFunc(url3, split))
}

func tryUmarshal() {
	type PWD struct {
		User string `json:"user"`
		PWD  string `json:"pwd"`
	}
	origin := PWD{PWD: "123"}
	str := `{"user": "test_user"}`
	fmt.Printf("before unmarshaled: %#v\n", origin)
	json.Unmarshal([]byte(str), &origin)
	fmt.Printf("after unmarshaled: %#v\n", origin)

	str2 := `{"pwd": "456"}`
	fmt.Printf("before unmarshaled: %#v\n", origin)
	json.Unmarshal([]byte(str2), &origin)
	fmt.Printf("after unmarshaled: %#v\n", origin)

}

func lenStr() {
	ch := "中国"
	en := "ch"
	// len ch: 6, len en: 2
	fmt.Printf("len ch: %d, len en: %d\n", len(ch), len(en))
	// len b of ch: 6, len b of en: 2
	fmt.Printf("len b of ch: %d, len b of en: %d\n", len([]byte(ch)), len([]byte(en)))
	// len rune of ch: 2, len rune of en: 2
	fmt.Printf("len rune of ch: %d, len rune of en: %d", len([]rune(ch)), len([]rune(en)))
}

func tryIfElse() {
	if a := "abc"; a == "" {
		fmt.Println("a is nil")
	} else {
		// 打印 "abc"
		fmt.Println("a: ", a)
	}
	/*	if a := "abc"; a == "" {
			fmt.Println("a is nil")
			b := "b"
		} else {
			// 报错，b 未定义
			fmt.Printf("a: %s, b: %s\b", a, b)
		}
	*/
}

func maxInt64() {
	mxInt64 := strconv.Itoa(math.MaxInt64)
	P(len(mxInt64), mxInt64, math.MaxInt64)
	mxInt32 := strconv.Itoa(math.MaxUint32)
	P(len(mxInt32), mxInt32)
}

// // 都会执行
// func init() {
// 	P("this is init at main package of 1")
// }
// func init() {
// 	P("this is init at main package of 2")
// }
// func init() {
// 	P("this is init at main package of 3")
// }

func tryConvertPanicToError() {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Printf("recovered from panic, panic: %s", err)
		}
	}()
	panic("ops...")
}

func tryAddDate() {
	date := time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC)
	nextMonth := date.AddDate(0, 2, 0)
	Pf("date : %s and nextMont: %s\n", date, nextMonth)
	date31 := date.AddDate(0, 0, 1)
	nextMonth = date31.AddDate(0, 2, 0)
	Pf("date : %s and nextMont: %s\n", date31, nextMonth)
}

func formatFloat() {
	f := 1381.15
	P("origin f: ", f)
	Pf("fmt format: %f\n", f)
	P("use strconv...")
	Pf("pre=2: %s\n", strconv.FormatFloat(f, 'f', 2, 32))
	Pf("pre=2, float64: %s\n", strconv.FormatFloat(f, 'f', 2, 64))
	var NIlInt []int
	P(NIlInt == nil, len(NIlInt))
	NIlInt = nil
	P(NIlInt == nil, len(NIlInt))
}

func tryCkecBankCard() {
	strs := []string{
		"1234567890123456",
		"123456789012345q",
		"123456789012345",
		"1234567890123454444",
		"12345678901234544",
		"12345678901234544222",
		"12345678901234544222s",
	}
	for _, v := range strs {
		P(checkBankCard(v))
	}
}

func checkBankCard(bankCard string) bool {
	pattern := `^(\d{16}|\d{19})$`
	ok, _ := regexp.MatchString(pattern, bankCard)
	return ok
}

func chanPrac() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Printf("the %dth v from pipe: %d \n", i, v)
	}
}
func chanPracBuffer() {
	ch := make(chan int, 11)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Printf("the %dth v from pipe: %d \n", i, v)
	}
	P("done")
}

// 100个id，对每个id做一系列操作，并返回结果（对结果的顺序没有要求）
// 限制启动goruntine个数为10
func graceGoruntine() {
	const (
		ID_NUM        = 100
		GORUNTINR_NUM = 10
	)
	result := make(chan string, GORUNTINR_NUM)

	work := func(id int, signal chan bool, result chan string) {
		str := strconv.Itoa(id)
		result <- str
		// 解除channel占用
		<-signal
	}

	// 分发，处理任务
	go func() {
		signal := make(chan bool, GORUNTINR_NUM)
		for i := 0; i < ID_NUM; i++ {
			signal <- true
			go work(i, signal, result)
		}
	}()
	// 处理返回的结果
	idStrs := make([]string, 0, ID_NUM)
	for i := 0; i < ID_NUM; i++ {
		v := <-result
		idStrs = append(idStrs, v)
	}
	fmt.Printf("the result: %+v\n", idStrs)
}

func makeMapLen() {
	m := make(map[string]string, 10)
	P(len(m))
	m["ka"] = "va"
	for k, v := range m {
		P(k, v)
	}
}

func tryFileJoin() {
	base1 := "ab/cd"
	base2 := "ab/cd/"
	base3 := "./ab/cd/"
	P(filepath.Join(base1, "e"))
	P(filepath.Join(base2, "e"))
	P(filepath.Join(base2, "/e"))
	P(filepath.Join(base3, "/e"))
}

func tryUpper() {
	str := "f9cbfc3f"
	P(strings.ToUpper(str))
	P(string(0x00), string(0x36), string(0x5c))
	P(string(byte(0x00)), string(byte(0x36)), string(byte(0x5c)))
	P(len("中国工商银行股份有限公司北京通州支行新华分理处"))
}

func tryLenStr() {
	strA := "abc"
	strB := "a中国"
	P("len A and B: ", len(strA), len(strB), len([]byte(strA)), len([]byte(strB)))
}

type A struct {
	B string
	c int
}

// func (a *A) String() string {
// return fmt.Sprintf("%#v", a)
// }

func tryPrintf() {
	sa := []*A{
		&A{"", 4},
		&A{"hello", 3},
		&A{"af", 0},
	}
	sb := []A{
		A{"adfa", 0},
	}
	fmt.Printf("\nvalue: %#v\n", sa[0])
	fmt.Printf("value+ : %+v, \nvalue# %#v,\nv: %v", sa, sa, sa)
	fmt.Printf("\n\nvalue+: %+v \nvalue#: %#v", sb, sb)
}

func tiny() {
	m := make(map[string]int)
	m["aa"]++
	P(m)
}

func tryCrypto() {
	pwd := "pwd"
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		P("bcrypt err: ", err)
	}

	P("first, pwd and hansh: \n", pwd, string(hash), len(string(hash)))
	hash2, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	P("second, pwd and hash: \n", pwd, string(hash2), len(string(hash2)))

	P("check pwd..")
	P("check hash1: ")
	err = bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	P(err == nil)
	err = bcrypt.CompareHashAndPassword(hash, []byte("pwds"))
	P(err == nil)
	P("check has2:")
	P("hash1 != hash2: ", string(hash) != string(hash2))
	err = bcrypt.CompareHashAndPassword(hash2, []byte(pwd))
	P(err == nil)
	u := uuid.New()
	P("uuid: ", u, len(u), len(uuid.New()), len(uuid.New()))
	unix := time.Now().Unix()
	unixStr := fmt.Sprintf("%d", unix)
	P("time: ", unix, len(unixStr))

}

func tryEmpty() {
	type s struct {
		FieldOne   int    `json:"one,omitempty"`
		FieldTwo   string `json:"two, omitempty"`
		FieldThree int    `json:"three"`
	}
	s1 := s{FieldTwo: ""}
	b1, _ := json.Marshal(s1)
	P("s1: ", string(b1))
}

func tryPrintfV() {
	testStruct := struct {
		field1 string
		filed2 int
		Field3 string
		Field4 bool `json:"filed4"`
	}{"filed1", 2, "filed3", true}
	fmt.Printf("testStruct: %v\n", testStruct)
	fmt.Printf("testStruct-pionter: %v\n", &testStruct)
}

func tryBuffer() {
	buf := bytes.NewBufferString("this is a buffer !")
	dst := make([]byte, 0, 10)
	dst2 := make([]byte, 10)
	dst3 := new([]byte)
	var dst4 []byte
	P(dst == nil, dst2 == nil, dst3 == nil, *dst3 == nil, dst4 == nil)
	n, err := buf.Read(dst)
	n2, err2 := buf.Read(dst2)
	n3, err3 := buf.Read(*dst3)
	P("dst: ", string(dst), "len: ", len(dst))
	P("dst2: ", string(dst2), "len: ", len(dst2))
	P("dst3: ", string(*dst3), "len: ", len(*dst3))
	P(n, err)
	P(n2, err2)
	P(n3, err3)
}

func testDefer() {
	start := time.Now()
	startCopy := start
	P("start: ", start)
	defer func(startCopy time.Time) {
		P("start at defer: ", start)
		P("startCopy at defer: ", startCopy)
	}(startCopy)
	time.Sleep(time.Second * 1)
	start = time.Now()
	startCopy = time.Now()
}

func WirdTest() {
	const Enone, Eio, Einval = 5, 2, 1

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	P("a")
	for i, v := range a {
		fmt.Printf("i: %d, v: %q\n", i, v)
	}
	P("s")
	for i, v := range s {
		fmt.Printf("i: %d, v: %q\n", i, v)
	}
	P("m")
	for i, v := range m {
		fmt.Printf("i: %d, v: %q\n", i, v)
	}
}

func fakeRequest() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	// test_uid := "45529"
	host := "0.0.0.0"
	port := "12231"
	dirver_trail_json := `[{"lat":"33.33066","lng":"121.284148","t":1472338663}]`
	driver_trail_form := url.Values{
		"session_id": []string{"test_uid"},
		"json":       []string{dirver_trail_json},
		"city":       []string{"上海"},
	}

	ticker := time.NewTicker(time.Second * 20)
	for t := range ticker.C {
		resp, err := http.PostForm(fmt.Sprintf("http://%s:%s/driver/trail", host, port), driver_trail_form)
		if err != nil {
			glog.Errorf("at %s, PostForm-err: %s\n", t, err)
		}
		defer resp.Body.Close()
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Errorf("at %s, ioutil.ReadAll-err: %s\n", t, err)
		}
		glog.Infof("at %s, fake post, fadback is %s", t, string(resp_body))
	}
}

func testMapLen() {
	dict := make(map[string]int)
	P("length of dict: ", len(dict))
	dictWithCap := make(map[string]int, 3)
	P("length of dictWithCap: ", len(dictWithCap))
	P("dict and dictWithCap: ", dict, dictWithCap)
	dictWithCap["a"] = 3
	P("length of dictWithCap: ", len(dictWithCap))
	P("dict and dictWithCap: ", dict, dictWithCap)

}

func tesArrayEqual() {
	var foo = [3]int{2, 3, 1}
	var bar = [3]int{2, 3, 4}
	var fish = [3]int{2, 3, 4}
	P("foo == bar: ", foo == bar)
	P("bar == fish: ", bar == fish)
}

func tesErrorFmt() {
	err := fmt.Errorf("a error is %s", "blabla...")
	glog.Infof("log the error: %s, use v: %v", err, err)
}

func httpRequest() {
	resp, err := http.Get(`http://115.29.34.206:12236/driver/trail?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`)
	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println(err, string(b))

	resp, err = http.Post(`http://115.29.34.206:12236/driver/trail?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`, "", nil)
	b, err = ioutil.ReadAll(resp.Body)

	fmt.Println(err, string(b))

	resp, err = http.Get(`http://115.29.34.206:12236/driver/orderTrail?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`)
	b, err = ioutil.ReadAll(resp.Body)

	fmt.Println(err, string(b))

	resp, err = http.Post(`http://115.29.34.206:12236/driver/orderTrail?&session_id=testpre&json=[{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0},{%22horAcc%22:10,%22t%22:1456733420,%22speed%22:0,%22isGps%22:1,%22order%22:%22%22,%22course%22:-1,%22lng%22:%22116.485900%22,%22lat%22:%2239.908870%22,%22type%22:0}]`, "", nil)
	b, err = ioutil.ReadAll(resp.Body)

	fmt.Println(err, string(b))
}

func tesSlice() {
	sl := []string{}
	asl := append(sl, "one")
	fmt.Println(asl, sl)
	sll := make([]string, 1, 4)
	asll := append(sll, "two")
	fmt.Println(asll, sll)
	// dup1()
	// charCount()
	tesFor()
}

// never found a better parterner for a long time, maybe met but I'am not brelliant to deserve it at that time. I need to focus and try to make me better and find a parterner to encourage and take care of each other

// review this
func charCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			glog.Infoln("is io.EOF")
			break
		}
		if line, prefix, _ := in.ReadLine(); string(line) == "" || string(line) == "\\n" {
			glog.Infof("line is %q, input should over, prefix: %t", string(line), prefix)
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func tesFor() {
	slice := []string{"abd", "dsad", "abc", "ddd"}
	for v := range slice {
		glog.Infoln(v)
	}
	for i, v := range slice {
		glog.Infoln(i, "\t", v)
	}
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		glog.Infoln("one scan, and the input is : ", input.Text())
		if input.Text() == "\n" {
			glog.Infoln("input is `\\n` over")
			break
		}
		if input.Text() == "" {
			glog.Infoln("the input is \"\", over")
			break
		}
		counts[input.Text()] += 1
	}

	for key, v := range counts {
		if v > 1 {
			fmt.Printf("%s:\t%d\n", key, v)
		}
	}
}

func tesMap() {
	var dic = map[string]int{}
	dic["a"] = 9
	delete(dic, "b")
	dic["c"] += 2
	fmt.Println(dic)

	// for k := range dic {
	// 	fmt.Println("K: ", k)
	// }
	var dic2 map[string]int
	var dic3 = map[string]int{}
	fmt.Println(len(dic2), len(dic3))
	fmt.Println("lookup: ", dic2["bb"], "-", dic3["bb"])
	delete(dic2, "kk")
	delete(dic3, "bb")
	fmt.Println("loop dic2: ")
	for k := range dic2 {
		fmt.Println("key in dic2: ", k)
	}
	for k := range dic3 {
		fmt.Println("key in dic3: ", k)
	}
	fmt.Println("==nil: ", dic2 == nil, dic3 == nil)
	// dic2["b"] = 8
	// fmt.Println(dic2)
}

func getEnv() {
	gopath := os.Getenv("ABC")
	fmt.Println("ABC: ", gopath)
}

func tryCall() {
	glog.Infoln("tryCall...")
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		glog.Infof("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
}

func timeAdd() {
	one := time.Now()
	glog.Infoln("one: ", one)
	two := one.Add(time.Minute * 10)
	glog.Infoln("two: ", two)
	glog.Infoln("one: ", one)

}

func closure() {
	closure1()
	closure2()
}

func closure1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i) // 555555
		}()
	}
	wg.Wait()
}

func closure2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			fmt.Println(j) // 01234(顺序可能会变)
		}(i)
	}
	wg.Wait()
}

// go run main.go -logtostderr
func syncRuntime() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://dict.youdao.com/w/currency/#keyfrom=dict2.top",
		"https://docs.mongodb.com/manual/mongo/",
		"http://www.runoob.com/mongodb/mongodb-q,uery.html",
		"http://studygolang.com/articles/2059",
	}
	glog.Infoln("fetching url..")
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			glog.Infoln("fetch url: ", url)
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			r, _ := http.Get(url)
			glog.Infoln("status: %s, code: %d, url is %s", r.Status, r.StatusCode, url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

// test loop
type MyInt int

func (m MyInt) String() string {
	return fmt.Sprint(int(m))
}

// panic
func intLoop() {
	var m MyInt = 888
	s := m.String()
	fmt.Println(s)
}

func build() {
	uids := []int{1256, 1457, 2799, 9448, 27078, 27090, 27095, 27104, 27148, 27160, 27988}
	// AND creator=%d;
	sql := "SELECT id FROM eventtopic WHERE `delete`=0"
	for _, uid := range uids {
		sql = fmt.Sprintf("%s%s", sql, fmt.Sprintf(" OR creator=%d", uid))
	}
	fmt.Println("sql: \n", sql+";")

}

func logFlush() {
	// go run main.go -alsologtostderr -log_dir="./"
	logDir := flag.Lookup("log_dir")
	testFlag := flag.Lookup("log_dir")
	glog.Infoln("lookup before parse", logDir.Name, logDir.Value, testFlag)
	err := flag.Set("log_dir", "test_value")
	glog.Errorln("err: ", err)
	glog.Infoln("abc..")
	glog.Infof("abc..%d", 123)
	logDir = flag.Lookup("log_dir")
	testFlag = flag.Lookup("log_dir")
	glog.Infoln("lookup before parse", logDir.Name, logDir.Value, testFlag.Name, testFlag.Value)
	glog.Flush()
}

func runes() {
	str := "hello中国"
	conv := []rune(str)
	P(str)
	for _, v := range conv {
		P(fmt.Sprintf("%d", v))
		P(string(v))
	}
	ss := []rune(str)[0]
	P(string(ss))
}

func convertInterface() {
	foo := Interface()
	tryOne, ok := foo.(map[string]string)
	fmt.Println(foo, tryOne, ok)
	tryTwo, ok := foo.(map[string]interface{})
	fmt.Println(tryTwo, ok)
}

func Interface() interface{} {
	foo := make(map[string]interface{})
	foo["key"] = "value"
	return foo
}

func makeNew() {

}

func crawl() {
	P("begin...")
	url := "http://www.applysquare.com/cn/"
	doc, err := goquery.NewDocument(url)
	P(doc, err)
}

func sorttt() {
	slice := []string{"a", "c", "b"}
	slice2 := []string{"a", "c", "b", "d"}

	sort.Strings(slice)
	sort.Strings(slice2)
	P(slice)
	P(slice2)
}

func loopMap() {
	dict := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}
	for key, val := range dict {
		P(key, val)
	}
}

func urlParse() {
	backend, err := url.Parse("http://localhost:8888/dir1/dir2/tail?q=123&q2=abc#alt")
	P(backend, err)
	P(MarshalJSONOrDie(backend))
}

func osPath() {
	d, err := os.Getwd()
	// 得到运行改命令时所在的目录, 比如在上一级目录运行该命令, 得到的目录就是上一级的
	P(d, err)
	dir1 := "/Users/mozhata/work/src/practice"
	dir2 := "/Users/mozhata/work/src/practice/"
	// /Users/mozhata/work/src
	P(path.Dir(dir1))
	// /Users/mozhata/work/src/practice
	P(path.Dir(dir2))
}

func buldCountySlug(slug string) string {
	countryKey := strings.Split(slug, ".")[0]
	if countryKey == "" {
		countryKey = "us"
	}
	return fmt.Sprintf("country_%s", countryKey)
}
func saveToSlice(source string, dest []string) {
	dest = append(dest, source)
}
func MarshalJSONOrDie(v interface{}) string {
	b, err := json.Marshal(v)
	Check(err)
	return string(b)
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
map-reduce
*/
/*
(doc) ->
	 return if doc.doc_type isnt 'Profile'
	 emit [doc.app_info.device_identifier.substr(0,2),doc.app_info.device_identifier.substr(0,4) ],null if doc.app_info.device_identifier
*/
// func Tt(s1,s2 string){
// 	print s1,s2
// }
func hello() {
	fmt.Println("hello world")
}
