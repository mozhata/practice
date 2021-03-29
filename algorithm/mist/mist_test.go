package mist_test

import (
	"fmt"
	"practice/algorithm/mist"
	"testing"
)

func TestMist(t *testing.T) {
	mist := mist.NewMist()
	for i := 0; i < 10; i++ {
		id := mist.Generate()
		fmt.Printf("i: %d, id: %d, id in byte: %064b\n", i, id, id)
	}
}
