package chap12

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func BasicRefect() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", 3)

	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("v: %v, x %v i: %v, v-d: %d\n", v, x, i, i)
}

func formatAtom(v reflect.Value) string {
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
		return v.Type().String() + "0x" + strconv.FormatUint(v.Pointer(), 16)
	default:
		return v.Type().String() + " value"
	}
}
