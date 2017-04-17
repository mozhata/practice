package trygob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
)

// BaiscGOB try basic gob usage:
// Create an encoder, transmit some values, receive them with a decoder.
func BaiscGOB() {
	type P struct {
		X, Y, Z int
		Name    string
	}
	type Q struct {
		Name string
		X, C int
	}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	var dataSent = P{1, 2, 3, "instance of P"}
	err := enc.Encode(dataSent)
	check(err)
	var dataReceived Q
	err = dec.Decode(&dataReceived)
	check(err)
	fmt.Printf("dataSent: %#v\tdataReceived: %#v", dataSent, dataReceived)
}

// InterfaceEncDec shows how to encode an interface value. The key distinction(主要区别) from regular types is to register the concrete type that implements the interface.
func InterfaceEncDec() {
	var network bytes.Buffer // Stand-in for the network.

	// We must register the concrete type for the encoder and decoder (which would
	// normally be on a separate machine from the encoder). On each end, this tells the
	// engine which concrete type is being sent that implements the interface.
	gob.Register(Point{})
	enc := gob.NewEncoder(&network)
	for i := 1; i < 4; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}
	// create decoder and receive some value
	dec := gob.NewDecoder(&network)
	for i := 1; i < 4; i++ {
		result := interfaceDecod(dec)
		fmt.Printf("result: %#v\tHypotenuse: %v\n", result, result.Hypotenuse())
	}
}

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// Pass pointer to interface so Encode sees (and hence sends) a value of
	// interface type. If we passed p directly it would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.
	err := enc.Encode(&p)
	check(err)
}
func interfaceDecod(dec *gob.Decoder) Pythagoras {
	var p Pythagoras
	err := dec.Decode(&p)
	check(err)
	return p
}

// EncodeDecode transmits a value that implements the custom encoding and decoding methods.
func EncodeDecode() {
	var network bytes.Buffer // Stand-in for a network connection
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{1, 2, 3})
	check(err)
	// create a decoder and receive value
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	check(err)
	fmt.Printf("received value: %#v\n", v)
	testMarshalUnmarshal()
}

// Vector type has unexported fields, which the package cannot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package. These interfaces are
// defined in the "encoding" package.
// We could equivalently use the locally defined GobEncode/GobDecoder
// interfaces.
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}
func (v *Vector) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}
func testMarshalUnmarshal() {
	source := Vector{2, 3, 4}
	b, err := source.MarshalBinary()
	check(err)
	var dest Vector
	err = dest.UnmarshalBinary(b)
	fmt.Printf("source: %#v\t, binaried: %q\n, umarshaled: %#v\n", source, b, dest)
}

func GobEncoderDecoder() {
	var network bytes.Buffer // Stand-in for a network connection
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Person{12, "test-name"})
	check(err)
	// create a decoder and receive value
	dec := gob.NewDecoder(&network)
	var v Person
	err = dec.Decode(&v)
	check(err)
	fmt.Printf("received value: %#v\n", v)
}

// Person impletemant GobEncoder and GobDecoder
type Person struct {
	age  int
	name string
}

func (p Person) GobEncode() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, p.age, p.name)
	return b.Bytes(), nil
}
func (p *Person) GobDecode(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &p.age, &p.name)
	return err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
