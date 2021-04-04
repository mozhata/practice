package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestSort(t *testing.T) {
	tables := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 1}, []int{1, 1, 2}},
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	}
	for _, cs := range tables {
		out := basic.BubbleSort(cs.input)
		Announce(t, out, assertions.ShouldResemble, cs.expect)
	}
}
