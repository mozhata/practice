package basic

import (
	"fmt"

	"git.100tal.com/jituan_AILab_Axer/scaffold/criteria/log"
)

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
*/
func TryDeDuplicate() {
	list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	fmt.Printf("before, list: %s\n", log.MarshalJSONOrDie(list))
	deDuplicated := deDuplicate(list)
	fmt.Printf("after, list: %#v\n", log.MarshalJSONOrDie(deDuplicated))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PushItem(val int) {
	head := l
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val: val}
}

func BuildList(arr []int) *ListNode {
	dumpy := &ListNode{}
	for _, i := range arr {
		dumpy.PushItem(i)
	}
	return dumpy.Next
}

/*
// TODO: 还需要重新刷一遍
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
*/
func Deduplicate(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}

func deDuplicate(list *ListNode) *ListNode {
	// 当前节点
	cur := list
	for cur != nil {
		// 有重复的删除
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		// 不重复, 则取下一个
		cur = cur.Next
	}
	return list
}

/*
TODO: 需要二刷的

https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
nil
1 -> 2 -> 2 -> 3
1 -> 1 -> 1 -> 3
1 -> 2 -> 3 -> 3
*/
func DelDuplicateElem(head *ListNode) *ListNode {
	// 这部分下面的逻辑包含了
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	dumpy := &ListNode{Next: head}
	head = dumpy
	// head 作为游标
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			// 从next开始, 删掉所有包含此值的节点
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			// 游标下移
			head = head.Next
		}
	}
	return dumpy.Next
}

/*
https://leetcode-cn.com/problems/reverse-linked-list/
*/
func ReverseLink(head *ListNode) *ListNode {
	var new *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = new
		new = head
		head = tmp
	}
	return new
}

// TODO: 不会, 先跳过
func ReverseBetween(head *ListNode, mIdx, nIdx int) *ListNode {
	// 假设m, count 超过限制会panic
	if head == nil {
		return nil
	}
	cur := head
	idx := 0
	for cur != nil {
		if idx == mIdx {
			// 开始反转

		}
		if idx == nIdx {
			// 反转完成, 收尾
		}

		// 指针下移
		cur = cur.Next
		idx++
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	// solution1: 递归

	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// next := head.Next
	// head.Next = swapPairs(next.Next)
	// next.Next = head
	// return next

	// solution2: 非递归
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		a, b := tmp.Next, tmp.Next.Next
		tmp.Next = b
		ttmp := b.Next
		b.Next = a
		a.Next = ttmp
		tmp = a
	}
	return pre.Next
}
