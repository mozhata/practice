package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestTravleTree(t *testing.T) {
	tables := []struct {
		input  *basic.TreeNode
		expect []int
	}{
		{
			nil,
			nil,
		},
		{
			&basic.TreeNode{Val: 1},
			[]int{1},
		},
		{
			&basic.TreeNode{Val: 1,
				Left: &basic.TreeNode{Val: 2,
					Left:  &basic.TreeNode{Val: 4},
					Right: &basic.TreeNode{Val: 5},
				},
				Right: &basic.TreeNode{Val: 3,
					Left:  &basic.TreeNode{Val: 6},
					Right: &basic.TreeNode{Val: 7},
				},
			},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			&basic.TreeNode{Val: 1,
				Left: &basic.TreeNode{Val: 2,
					Left:  &basic.TreeNode{Val: 4},
					Right: &basic.TreeNode{Val: 5},
				},
				Right: &basic.TreeNode{Val: 3,
					Right: &basic.TreeNode{Val: 7},
				},
			},
			[]int{1, 2, 4, 5, 3, 7},
		},
	}
	for _, cs := range tables {
		out := basic.PreOrderTravleTree(cs.input)
		Announce(t, out, assertions.ShouldResemble, cs.expect)
	}
}
