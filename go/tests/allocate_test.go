package tests

import "testing"

func allocate(n int) {
	if n < 5 {
		return
	}
	slice := make([]string, 10000)
	_ = slice
}

func BenchmarkLess5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allocate(4)
	}
}

func BenchmarkSurpass5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allocate(7)
	}
}
