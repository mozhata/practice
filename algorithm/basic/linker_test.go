package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestDelDuplicateElem(t *testing.T) {
	rawTable := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{1, 2, 2, 3}, []int{1, 3}},
		{[]int{1, 1, 1, 3}, []int{3}},
		{[]int{1, 2, 3, 4}, []int{1, 2}},
	}
	tables := []struct {
		Input  *basic.ListNode
		Expect *basic.ListNode
	}{}
	inh := basic.ListNode{}
	exh := basic.ListNode{}
	for _, ra := range rawTable {
		for _, i := range ra.input {
			inh.Next = &basic.ListNode{Val: i}
		}
		for _, j := range ra.expect {
			exh.Next = &basic.ListNode{Val: j}
		}
	}
	for _, i := range tables {
		o := basic.DelDuplicateElem(i.Input)
		Announce(t, o, assertions.ShouldResemble, i.Expect)
	}
}
