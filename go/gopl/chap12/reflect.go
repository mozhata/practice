package chap12

import (
	"fmt"
	"io"
	"os"
	"reflect"
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

	}
}
