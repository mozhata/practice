package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestDelDuplicateElem(t *testing.T) {
	tables := []struct {
		Input  *basic.ListNode
		Expect *basic.ListNode
	}{
		{
			Input:  nil,
			Expect: nil,
		},
		{
			Input: &basic.ListNode{
				Val: 1,
				Next: &basic.ListNode{
					Val: 2,
					Next: &basic.ListNode{
						Val: 2,
						Next: &basic.ListNode{
							Val: 3,
						},
					},
				},
			},
			Expect: &basic.ListNode{
				Val: 1,
				Next: &basic.ListNode{
					Val: 3,
				},
			},
		},
		{
			Input: &basic.ListNode{
				Val: 1,
				Next: &basic.ListNode{
					Val: 1,
					Next: &basic.ListNode{
						Val: 1,
						Next: &basic.ListNode{
							Val: 3,
						},
					},
				},
			},
			Expect: &basic.ListNode{
				Val: 3,
			},
		},
		{
			Input: &basic.ListNode{
				Val: 1,
				Next: &basic.ListNode{
					Val: 2,
					Next: &basic.ListNode{
						Val: 3,
						Next: &basic.ListNode{
							Val: 3,
						},
					},
				},
			},
			Expect: &basic.ListNode{
				Val: 1,
				Next: &basic.ListNode{
					Val: 2,
				},
			},
		},
	}
	for _, i := range tables {
		o := basic.DelDuplicateElem(i.Input)
		Announce(t, o, assertions.ShouldResemble, i.Expect)
	}
}

/*

func TestSetLen(t *testing.T) {
	s := basic.IntSet{}
	Announce(t, s.Len(), assertions.ShouldEqual, 0)
	s.Add(0)
	Announce(t, s.Len(), assertions.ShouldEqual, 1)
	s.Add(2)
	Announce(t, s.Len(), assertions.ShouldEqual, 2)
	s.Add(3)
	Announce(t, s.Len(), assertions.ShouldEqual, 3)
	s.AddAll(4, 5)
	Announce(t, s.Len(), assertions.ShouldEqual, 5)

}
*/
