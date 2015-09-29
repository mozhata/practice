package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/AlexRaylight/go-simplejson"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {

	// TrySimpleJson()
	// tryComma()
	// tryFilePath()
	// TryRegex()
	// var a []string
	// P(a == nil)
	// tryRange()

	// josnRawMessage()
	// tryMap()
	// BBase64()
	// tryBuffer()
	// Or("")
	// r := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Println("Enter command->")
	// 	line, _, _ := r.ReadLine()
	// 	fmt.Println(string(line))
	// }

}
func tryRange() {
	// 若range能写成for index, val := range slice{...}
	// 的形式,若写成for index := range slice {...} 则只遍历索引
	// Map 则只遍历key
	slice := []string{"a", "12", "33"}
	for i, v := range slice {
		P(i, v)
	}
	for s := range slice {
		P(s)
	}

	sliceB := []byte{'a', 'b', 'c'}
	for i, v := range sliceB {
		P(i, v)
	}
	for s := range sliceB {
		P(s)
	}

	dic := map[string]interface{}{"a": 1, "key1": 2, "slice": []string{"abs", "dada~"}}
	for key, val := range dic {
		P(key, val)
	}
	for key := range dic {
		P(key)
	}
}
func tryComma() {
	var bt interface{}
	bt = []byte("a b c")
	// 被判断的类型必须是interface{}
	// s 先被声明为对应的变量(string),
	// 若ok为true,再给s赋bt的值值,(若为false,则s为零值)
	s, ok := bt.(string)
	b, ok := bt.([]byte)
	P(s, ok, s == "")
	P(b, ok)
}

func TrySimpleJson() {
	body := []byte(jsonString)
	json, _ := simplejson.NewJSON(body)
	P(json.Get("data", "content", "body").String())
	content := json.Get("data", "content")
	P("content on *json-type: ", content)
	byte_data, _ := json.Get("data", "content").MarshalJSON()
	// P("not marshaled, just use string(): ", string(json.Get("data", "content")))		// failed
	P("content MarshalJsoned: ", string(byte_data))
	// not follow at archture:
	P("not follow archery: ", json.Get("title"))

	json2_byte := []byte(json2)
	json, _ = simplejson.NewJSON(json2_byte)
	// for i, v := range json.Get("conversations").Array() {
	//  P(i, v)
	// }
	qa_activity := json.Get("conversations").JSONArray()[0].Get("data", "qa_activity")
	qa_activity1, _ := qa_activity.MarshalJSON()
	P("stringed qa_activity", string(qa_activity1))
	P("action: ", qa_activity.Get("action").String())
	P("num_unread: ", qa_activity.Get("num_unread").Int())
	P("thread_id: ", qa_activity.Get("thread_id").String())
	data_raw := json.Get("conversations").JSONArray()[0].Get("data_raw").String()
	P("raw_data_raw: ", data_raw)
}

var jsonString string = `{
   "data": {
       "author": {
           "id": "909e7af2-020e-4f11-a637-d77206f72903"
       },
       "content": {
           "body": "<p>问题一说明</p>"
       },
       "created": "2015-07-31T07:34:36",
       "id": "8MRAsh6VC",
       "modified": "2015-08-14T08:08:32",
       "next_id": "3",
       "replies": [
           {
               "author": {
                   "id": "017d49c9-2646-4cc4-b0fc-47d70a600abb"
               },
               "content": {
                   "body": "<p>dfhaderfhe yh5是个</p>"
               },
               "created": "2015-08-04T03:27:31",
               "id": "1",
               "modified": "2015-08-14T07:47:58"
           },
           {
               "author": {
                   "id": "cb05b384-6a4a-47ca-bf96-093c867c2764"
               },
               "content": {
                   "body": "<p>four</p>"
               },
               "created": "2015-08-14T07:56:06",
               "id": "2",
               "modified": "2015-08-14T07:56:06"
           }
       ],
       "tag": {
           "category": {
               "key": "uncategorized"
           },
           "tags": [
               {
                   "key": "选择困难症"
               }
           ]
       },
       "title": "问题1--edited by two"
   }
}`
var json2 string = `{
    "conversations": [
        {
            "created": "2012-11-01T23:15:42.123456Z",
            "data": {
                "qa_activity": {
                    "action": "new_tag_thread",
                    "num_unread": 1,
                    "thread_id": "中文",
                    "thread_title": "中文"
                }
            },
            "data_raw": "eyJxYV9hY3Rpdml0eSI6eyJhY3Rpb24iOiJuZXdfdGFnX3RocmVhZCIsIm51bV91bnJlYWQiOjEsInRocmVhZF9pZCI6IuS4reaWhyIsInRocmVhZF90aXRsZSI6IuS4reaWhyJ9fQ==",
            "from_user": "999",
            "id": 23,
            "num_messages": 2,
            "param0": "中文",
            "param1": "new_tag_thread",
            "read": true,
            "to_user": "100",
            "with_user": "999"
        }
    ]
}`

func tryBuffer() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprint(&b, "world, ...")
	b.WriteTo(os.Stdout)
}
func Or(host string) {
	if host == "" {
		host = "123"
	}
	P(host)
}
func BBase64() {
	str := "eyJxYV9hY3Rpdml0eSI6eyJhY3Rpb24iOiJuZXdfdGFnX3RocmVhZCIsIm51bV91bnJlYWQiOjIsInRocmVhZF9pZCI6InRhZ2EiLCJ0aHJlYWRfdGl0bGUiOiJ0YWdhIn19"
	data, _ := base64.StdEncoding.DecodeString(str)
	P(string(data))
}

func josnRawMessage() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
    {"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
    {"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
  ]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	CheckErr(err)

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		CheckErr(err)

		P(c.Space, dst)
	}
}
func trySetPath() {

}

// map遍历 可以只遍历key
func tryMap() {
	dic := map[string]interface{}{"a": 1, "key1": 2, "slice": []string{"abs", "dada~"}}
	P(dic)
	keys := []string{}

	// 可以只遍历key
	for key := range dic {
		keys = append(keys, key)
	}

	P(keys)
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func tryFilePath() {
	path := "/home/zyk/test/sublime_imfix.c"
	P(filepath.Ext(path)) // 	.c
	path2 := "/home/zyk/test/"
	path3 := "/home/zyk/test"
	P(filepath.Base(path), filepath.Base(path2), filepath.Base(path3)) // sublime_imfix.c test test
}
func TryRegex() {
	text := `allText
	Contact
	  Alecia Denmark, Director of Undergraduate Admissions
	  800-548-7638
	  email the program
	  visit the website`
	// match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	r, _ := regexp.Compile("[0-9-]{12}")

	fmt.Println(r.FindString(text))
}
