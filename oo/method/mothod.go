package method

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type Rect struct {
	x, y          float64
	width, height float64
}
type Foo struct {
	Rect
	name string
}
type techer struct {
	human
	Name string
	Age  int64
}
type student struct {
	human
	Name string
	Age  int64
}
type human struct {
	Sex int
}
type A struct {
	B
	C
}
type B struct{ Name string }
type C struct{ Name string }

func Typ() {
	a := &A{B: B{Name: "B"}, C: C{Name: "C"}}
	fmt.Println(a.B.Name)

}

type Str []string

func (s *Str) Find(name string) string {
	if len(s) == 0 {
		return "empty slice"
	}
	for _, str := range s {
		if str == name {
			return "found"
		}
	}
	return "not found"
}

func Pe() {
	a := &techer{Name: "joe", Age: 33, human: human{Sex: 8}}
	b := &student{Name: "joe", Age: 33, human: human{Sex: 9}}
	a.Name = "joe2"
	a.Age = 12
	a.Sex = 88
	b.human.Sex = 999
	fmt.Println(*a, "\n", *b)
}
func (r *Rect) Area() float64 {
	return r.width * r.height
}
func (r *Foo) SayHi() {
	fmt.Println("hello,I am type of Foo,name is: ", r.name)
}
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

// test
var a Integer = 3
var rect1 = &Rect{2, 3, 4, 5}
var rect2 = &Foo{Rect{0, 0, 9, 5}, "puppy"}

func Tarea() {
	fmt.Println(rect1.Area())
}
func Tinherit() {
	rect2.SayHi()
	fmt.Println(rect2.Rect.Area())
	fmt.Println(rect2.Area())
}
func Tless() {
	if a.Less(4) {
		fmt.Println(a, "Less 4")
	} else {
		fmt.Println(a, "not less than 4")
	}
}
func Tadd() {
	a.Add(2)
	fmt.Println("a: ", a)
}
