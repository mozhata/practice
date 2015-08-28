package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/AlexRaylight/go-simplejson"
	"github.com/golang/glog"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// TrySimpleJson()
	var a []string
	P(a == nil)
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

func TrySimpleJson() {
	body := []byte(jsonString)
	json, _ := simplejson.NewJSON(body)
	P(json.Get("data", "content", "body").String())
	content := json.Get("data", "content")
	P("content on *json-type: ", content)
	byte_data, _ := json.Get("data", "content").MarshalJSON()
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

// 通过一个string调用对应的函数的方法:
func runFunction(function string) error {
	glog.Infof("Running func: %v", function)
	defer glog.Infof("Finished func: %v", function)

	funcMap := map[string]func() error{
		"news":      funcA, // func funcA() error
		"professor": funcB, // func funcB() error
		"program": func() error {
			funcC(nil, true) // func serveProgram(r *mux.Router, runOnce bool)
			return nil
		},
	}

	f, ok := funcMap[function]
	if !ok {
		return fmt.Errorf("Unknown func name: %s", function)
	}
	return f()
}
