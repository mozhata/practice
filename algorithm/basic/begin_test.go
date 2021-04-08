package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestRotate(t *testing.T) {
	tables := []struct {
		input  [][]int
		expect [][]int
	}{
		{
			nil, nil,
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
	}
	for _, cs := range tables {
		basic.Rotate(cs.input)
		Announce(t, cs.input, assertions.ShouldResemble, cs.expect)
	}
}

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

func TestMinWindowV1(t *testing.T) {
	tables := []struct {
		input  []string
		expect string
	}{
		{
			[]string{"ADOBECODEBANC", "ABC"},
			"BANC",
		},
	}
	for _, cs := range tables {
		out := basic.MinWindowV1(cs.input[0], cs.input[1])
		Announce(t, out, assertions.ShouldEqual, cs.expect)
	}
}
func TestDecodeStr(t *testing.T) {
	tables := []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"a", "a"},
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
	}
	for _, cs := range tables {
		idx := basic.DecodeStr(cs.input)
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
