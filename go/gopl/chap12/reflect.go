package chap12

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var Logger = log.New(os.Stdout, "", log.Lshortfile)

func BasicRefect() {
	// t := reflect.TypeOf(3)
	// fmt.Println(t.String())

	// var w io.Writer = os.Stdout
	// fmt.Println(reflect.TypeOf(w))
	// fmt.Printf("%T\n", 3)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("value of 3: %v\n", v)
	fmt.Println(v.String())
	t := v.Type()
	Logger.Printf("v.Type: %s", t)
	sv := reflect.ValueOf("abc")
	Logger.Printf("sv.Type: %s", sv.Type())

	x := v.Interface()
	i := x.(int)
	fmt.Printf("v: %v, x %v i: %v, v-d: %d\n", v, x, i, i)
}

func TryAny() {
	var x int = 9
	var d time.Duration = time.Nanosecond * 1
	type mint int
	var m mint = 5
	Logger.Println(Any(x))
	Logger.Println(Any(d))
	Logger.Println(Any(m))

	Logger.Println(Any([]time.Duration{d}))
	Logger.Println(Any([]mint{m}))
	Logger.Println(Any([]int{x}))

	Logger.Println(Any(func() {}))
	Logger.Println(Any(time.Now()))
	Logger.Println(Any(struct{}{}))
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	Logger.Println("kind: ", v.Kind())
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".valud", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
