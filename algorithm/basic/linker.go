package basic

import (
	"fmt"

	"git.100tal.com/jituan_AILab_Axer/scaffold/criteria/log"
)

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
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
