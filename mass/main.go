package main

import (
	// . "bitbucket.org/applysquare/applysquare-go/pkg/discussion"

	"fmt"
	"strconv"
	"time"

	// "github.com/robfig/cron"
	// "github.com/stretchr/testify/assert"
	// "testing"
)

var P func(...interface{}) (int, error) = fmt.Println
var parse func(string, int, int) (int64, error) = strconv.ParseInt
var atoi func(string) (int, error) = strconv.Atoi
var EmailPattern = `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})$`

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
	// s1 := "abcdditui"
	// // s2 := "1234"
	// ss := md5.Sum([]byte(s1))
	// P("md5: ", ss)
	// fmt.Printf("%x", ss)
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
	var ss = struct {
		a int
		b string
	}{}
	P(ss)
	P(&ss == nil)

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
