package main

import "fmt"

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	//普通类型向接口类型的转换是隐式的。
	//接口类型向普通类型转换需要类型断言

	// 普通类型向接口类型的转换是隐式的,是编译期确定的
	var val interface{} = "hello"
	P(val)
	val = []byte{'a', 'b', 2}
	P(val)

	// 接口类型向普通类型转换有两种方式：Comma-ok断言和switch测试

	// Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，T是普通类型。
	type Html []interface{}
	html := make(Html, 4)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = 33

	for _, element := range html {
		if value, ok := element.(string); ok {
			P(value, "is string")
		} else if value, ok := element.([]byte); ok {
			P(value, "is []byte")
		}
	}

	// switch 测试:
	for _, element := range html {
		switch value := element.(type) {
		case string:
			P(value, "is string")
		case []byte:
			P(value, "is []byte")
		default:
			P("Unkown type")
		}
	}
	// Comma-ok断言还支持另一种简化使用的方式：value := element.(T)。但这种方式不建议使用，因为一旦element.(T)断言失败，则会产生运行时错误。
	tem := html[0].(string)
	P(tem)
	// tem2 will panic, 因为断言失败
	tem2 := html[0].(int)
	P(tem2)
}
