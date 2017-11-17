package util

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

// BuildErrMsg used to format error message
func BuildErrMsg(msgs ...interface{}) string {
	if len(msgs) == 0 || msgs == nil {
		return ""
	}
	if len(msgs) == 1 {
		if v, ok := msgs[0].(string); ok {
			return v
		}
		if v, ok := msgs[0].(error); ok {
			return v.Error()
		}
	}
	if len(msgs) > 1 {
		return fmt.Sprintf(msgs[0].(string), msgs[1:]...)
	}
	return ""
}

func Debug(a ...interface{}) {
	_, fileName, line, _ := runtime.Caller(1)
	index := strings.LastIndex(fileName, "practice/")
	if index > 0 {
		fileName = fileName[index+len("practice/"):]
	}
	msg := BuildErrMsg(a...)
	fmt.Printf("\n%s:%d\n%s\n", fileName, line, msg)
}

func MarshalJSONOrDie(v interface{}) string {
	b, e := json.MarshalIndent(v, "", "  ")
	if e != nil {
		panic(e)
	}
	return string(b)
}
