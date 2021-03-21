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

func DelDuplicateElem(head *ListNode) *ListNode {
	dumpy := ListNode{Next: head}
	pre := dumpy
	head = head
	var watchedRemove bool
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			head = head.Next.Next
			watchedRemove = true
		}
	}
	return nil
}
