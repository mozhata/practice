package basic

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序递归遍历: 略

// 前序非递归遍历
func PreOrderTravleTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var restult []int
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			restult = append(restult, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		//
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return restult
}

/*
中序非递归
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for len(stack) > 0 || root != nil {
		// 压栈到最左
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right

	}

	return result
}

// 后续非递归

// DFS(深度优先) 深度搜索-从上到下

// DFS 深度搜索-从下向上（分治法）

// BFS(广度优先)层次遍历
