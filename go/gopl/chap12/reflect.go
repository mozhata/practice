package chap12

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var Logger = log.New(os.Stdout, "", log.Lshortfile)

type Person struct {
	Name string
}
type Actor struct {
	Person
	Role string
	hid  int
}
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Actors          []Actor
	Oscars          []string
	Sequel          *string
}

var sequel = "this is sequel"
var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Actors: []Actor{
		Actor{
			Person: Person{
				Name: "person a",
			},
			Role: "role of person a",
		},
		Actor{
			Person: Person{
				Name: "person b",
			},
			Role: "role of person b",
			hid:  9,
		},
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
	Sequel: &sequel,
}

func TryChangeValue() {

}

func TrySMarshal() {
	b, err := Marshal(strangelove)
	if err != nil {
		panic(fmt.Sprintf("Marshal value to S code failed: %s", err.Error()))
	}
	Logger.Printf("S coded: \n%s\n", string(b))
}

func TryDisplay() {
	Display("strangelove", strangelove)
	// Display("os.Stderr", os.Stderr)
	// Display("rV", reflect.ValueOf(os.Stderr))

	func() {
		var i interface{} = 3
		Display("i", i)
		Display("&i", &i)
	}()

	/*
		func() {
			type Cycle struct {
				Value int
				Tail  *Cycle
			}
			var c Cycle
			c = Cycle{42, &c}
			Display("c", c)
		}()
	*/

	func() {
		type Cycle struct {
			Value int
			Tail  *Cycle
		}
		var c Cycle
		c = Cycle{42, &c}
		DisplayV2("c", c)
	}()
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

func Marshal(v interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	err := encode(&buf, reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func DisplayV2(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	count := 0
	displayV2(&count, 10, name, reflect.ValueOf(x))
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	// Logger.Println("kind: ", v.Kind())
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

func displayV2(count *int, top int, path string, v reflect.Value) {
	if *count > top {
		return
	}
	print := func(count *int, format string, a ...interface{}) {
		fmt.Printf(format, a...)
		*count = *count + 1
	}
	switch v.Kind() {
	case reflect.Invalid:
		print(count, "%s = invalid", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			displayV2(count, top, fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayV2(count, top, fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			displayV2(count, top, fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			print(count, "%s = nil\n", path)
		} else {
			displayV2(count, top, fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			print(count, "%s = nil\n", path)
		} else {
			print(count, "%s.type = %s\n", path, v.Elem().Type())
			displayV2(count, top, path+".valud", v.Elem())
		}
	default: // basic types, channels, funcs
		print(count, "%s = %s\n", path, formatAtom(v))
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

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
