package bit

import (
	"bytes"
	"fmt"
)

func TryBit() {
	var (
		i uint8 = 1
		j int8  = 1
	)
	fmt.Printf("i: %08b, j: %08b\n", i, j)
	i, j = i<<6, j<<6
	fmt.Printf("i: %08b, j: %08b\n", i, j)
	i, j = i<<1, j<<1
	fmt.Printf("i: %08b, j: %08b\n", i, j)
	i, j = i<<1, j<<1
	fmt.Printf("i: %08b, j: %08b\n", i, j)
	fmt.Printf("div: %v, mod: %v\n", -8/3, -8%3)
}

func TryBitmap() {
	bm := New()
	src := []uint64{4, 5, 2, 5555, 7, 888, 9999, 33, 999, 222323, 123, 555, 1, 9, 3, 2, 8, 3}
	for _, v := range src {
		bm.Add(v)
	}
	fmt.Printf("bm: %s", bm.String())
}

type Bitmap struct {
	words  []uint64
	length uint64
}

func New() *Bitmap {
	return &Bitmap{}
}

// Add add
func (b *Bitmap) Add(n uint64) {
	idx, bit := int(n/64), n%64
	if idx >= len(b.words) {
		b.words = append(b.words, make([]uint64, idx-len(b.words)+1)...)
	}
	if b.words[idx]&(1<<bit) == 0 {
		b.words[idx] = b.words[idx] | (1 << bit)
		b.length++
	}
}

// Has has
func (b *Bitmap) Has(n uint64) bool {
	idx, bit := int(n/64), n%64
	return idx < len(b.words) && b.words[idx]&(1<<bit) != 0
}

func (b *Bitmap) Len() uint64 {
	return b.length
}

func (b *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range b.words {
		if v == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if b.words[i]&(1<<j) != 0 {
				fmt.Fprintf(&buf, "%d ", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf, "\nlength: %d", b.length)

	return buf.String()
}
