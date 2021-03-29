package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestIndexOfSubStr(t *testing.T) {
	tables := []struct {
		input  []string
		expect int
	}{
		{[]string{"", ""}, 0},
		{[]string{"abcdef", ""}, 0},
		{[]string{"abccdef", "ab"}, 0},
		{[]string{"abccdef", "bc"}, 1},
		{[]string{"abccdef", "c"}, 2},
		{[]string{"abccdef", "def"}, 4},
		{[]string{"abccdef", "f"}, 6},
		{[]string{"abccdef", "fa"}, -1},
	}
	for _, cs := range tables {
		idx := basic.IndexOfSubStr(cs.input[0], cs.input[1])
		Announce(t, idx, assertions.ShouldEqual, cs.expect)
	}
}
func TestTryForI(t *testing.T) {
	tables := []struct {
		input  int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{111, 111},
	}
	for _, cs := range tables {
		out := basic.TryForI(cs.input)
		Announce(t, out, assertions.ShouldEqual, cs.expect)
	}
}
