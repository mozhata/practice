package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestBuildList(t *testing.T) {
	table := []struct {
		input  []int
		expect *basic.ListNode
	}{
		{
			input:  nil,
			expect: nil,
		},
		{
			input:  []int{1},
			expect: &basic.ListNode{Val: 1},
		},
		{
			input: []int{0, 1, 2},
			expect: &basic.ListNode{Val: 0,
				Next: &basic.ListNode{Val: 1,
					Next: &basic.ListNode{Val: 2}}},
		},
		{
			input: []int{2, 3, 5, 6, 8},
			expect: &basic.ListNode{Val: 2,
				Next: &basic.ListNode{Val: 3,
					Next: &basic.ListNode{Val: 5,
						Next: &basic.ListNode{Val: 6,
							Next: &basic.ListNode{Val: 8}}}}},
		},
	}
	for _, tt := range table {
		out := basic.BuildList(tt.input)
		Announce(t, out, assertions.ShouldResemble, tt.expect)
	}
}

func TestDeduplicate(t *testing.T) {
	rawTable := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 3}, []int{1, 3}},
		{[]int{1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}
	for _, ra := range rawTable {
		input := basic.BuildList(ra.input)
		expect := basic.BuildList(ra.expect)
		o := basic.Deduplicate(input)
		Announce(t, o, assertions.ShouldResemble, expect)
	}
}
func TestDelDuplicateElem(t *testing.T) {
	rawTable := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{1, 2, 2, 3}, []int{1, 3}},
		{[]int{1, 1, 1, 3}, []int{3}},
		{[]int{1, 2, 3, 3}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}
	for _, ra := range rawTable {
		input := basic.BuildList(ra.input)
		expect := basic.BuildList(ra.expect)
		o := basic.DelDuplicateElem(input)
		Announce(t, o, assertions.ShouldResemble, expect)
	}
}

func TestReverseLink(t *testing.T) {
	rawTable := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}
	for _, ra := range rawTable {
		input := basic.BuildList(ra.input)
		expect := basic.BuildList(ra.expect)
		o := basic.ReverseLink(input)
		Announce(t, o, assertions.ShouldResemble, expect)
	}
}
