package basic

import (
	"bytes"
	"fmt"
)

func TryIntSet() {
	set := IntSet{}
	set.AddAll(1, 2, 3)
	fmt.Printf("set: %v\n", set)
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	idx, bit := x/64, x%64
	return len(s.words) > idx && s.words[idx]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	idx, bit := x/64, x%64
	for idx >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[idx] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

func (s *IntSet) Remove(x int) {
	idx, bit := x/64, x%64
	if idx >= len(s.words) {
		return
	}
	s.words[idx] &^= 1 << bit
}

func (s *IntSet) UnionWith(t IntSet) {
	for idx, word := range t.words {
		if idx < len(s.words) {
			s.words[idx] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}
func (s *IntSet) DiffWith(t IntSet) {
	for idx, word := range t.words {
		if idx < len(s.words) {
			s.words[idx] ^= word
		}
	}
}
func (s *IntSet) SymmetricDiffWith(t IntSet) {}
func popCount(x uint64) int {
	var count int
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
func (s *IntSet) Len() int {
	var count int
	for _, w := range s.words {
		count += popCount(w)
	}
	return count
}
func (s *IntSet) Clear() {
	s.words = []uint64{}
}
func (s *IntSet) Copy() *IntSet {
	newWords := make([]uint64, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}
func (s *IntSet) Elems() []uint64 {
	var result []uint64
	for i, w := range s.words {
		for j := uint64(0); j < 64; j++ {
			if w&(1<<j) != 0 {
				result = append(result, uint64(i)*64+j)
			}
		}
	}
	return result
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	elms := s.Elems()
	if len(elms) > 0 {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		for _, e := range elms {
			fmt.Fprintf(&buf, "%d", e)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

/*
&      位运算 AND
|      位运算 OR
^      二元: 位运算 XOR (不相等得1, 相等得0)  | 一元: 按位取反
&^     位清空 (AND NOT)
<<     左移
>>     右移
*/
func TryBitSet() {
	var (
		x uint8 = 1<<1 | 1<<5
		y uint8 = 1<<1 | 1<<2
	)
	fmt.Printf("x    in byte: %08b, %d\n", x, x)       // 00100010 the set of {1,5}
	fmt.Printf("^x   in byte: %08b, %d\n", ^x, ^x)     // 11011101 按位取反
	fmt.Printf("x<<1 in byte: %08b, %d\n", x<<1, x<<1) // 01000100 the set of {2,6}
	fmt.Printf("x>>1 in byte: %08b, %d\n", x>>1, x>>1) // 00010001 the set of {0,4}
	fmt.Printf("y in byte: %08b, %d\n", y, y)          // 00000110 the set of {1,2}
	fmt.Printf("y<<1 in byte: %08b, %d\n", y<<1, y<<1) // 00001100 the set of {2,3}
	fmt.Printf("x|y  in byte: %08b, %d\n", x|y, x|y)   // 00100110 the set of {1,2,5}
	fmt.Printf("x&y  in byte: %08b, %d\n", x&y, x&y)   // 00000010 the set of {1}
	fmt.Printf("x^y  in byte: %08b, %d\n", x^y, x^y)   // 00100100 the set of {2,5}
	fmt.Printf("x&^y in byte: %08b, %d\n", x&^y, x&^y) // 00100100 the set of {2,5}
	// 通过&运算 取出集合中的元素
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Printf("item in set x: %d\n", i)
		}
	}
	for i := uint(0); i < 8; i++ {
		if x>>1&(1<<i) != 0 {
			fmt.Printf("%d in set x>>1\n", i)
		}
	}
	/* 	// 有符号整数位移运算 (有符号整数采用2的补码形式表示) 有符号整数的取反, 位移都较难理解, 慎用
	   	var z int8 = 1<<0 | 1<<6
	   	fmt.Printf("-1   in byte: %08b, %d\n", int8(-1), int8(-1))   // -0000001
	   	fmt.Printf("^-1  in byte: %08b, %d\n", ^int8(-1), ^int8(-1)) // 00000000
	   	fmt.Printf("z    in byte: %08b, %d\n", z, z)                 // 01000001
	   	fmt.Printf("^z   in byte: %08b, %d\n", ^z, ^z)               // -1000010, 因使用补码的缘故, -1000001 +1 -> -1000010
	   	fmt.Printf("z<<1 in byte: %08b, %d\n", z<<1, z<<1)           // -1111110
	   	fmt.Printf("z<<2 in byte: %08b, %d\n", z<<2, z<<2)
	   	fmt.Printf("z>>1 in byte: %08b, %d\n", z>>1, z>>1)
	*/
}
