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

/*
二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
递归法, 或者用广度优先
*/
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queque := make([]*TreeNode, 0)
	queque = append(queque, root)
	var layer int
	for len(queque) > 0 {
		num := len(queque)
		// 把这一层的都取出来
		for i := 0; i < num; i++ {
			item := queque[0]
			queque = queque[1:]
			// 把下一层的放到队列里面, 一次循环的时候处理
			if item.Left != nil {
				queque = append(queque, item.Left)
			}
			if item.Right != nil {
				queque = append(queque, item.Right)
			}
		}
		layer++
	}
	return layer
}

/*
二叉树的层序遍历
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/bfs-de-shi-yong-chang-jing-zong-jie-ceng-xu-bian-l/
*/
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 计算当前层的元素个数
		lLen := len(queue)
		for i := 0; i < lLen; i++ {
			// 取出一个元素, 放入list
			item := queue[0]
			queue = queue[1:]
			list = append(list, item.Val)
			// 把下一层的元素放入队列
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}

		result = append(result, list)
	}
	return result
}

/*
https://leetcode-cn.com/problems/as-far-from-land-as-possible/
*/

/*
柱状图中最大的矩形
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
*/
//  暴力解法
// TODO: 优化解暂时不会, 先放过
func LargestRetangleAreaForce(heights []int) int {
	var result int
	for k, v := range heights {
		// 寻找当前高度的左右边界
		left, right := k, k
		for left > 0 && heights[left-1] >= v {
			left--
		}
		for right < len(heights)-1 && heights[right+1] >= v {
			right++
		}
		result = max(result, (right-left+1)*v)
	}
	return result
}
