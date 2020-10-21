/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
*/
package leetcode

/*
解法: 双指针


结果: 执行用时 :4 ms 内存消耗 :3.1 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

func deleteDuplicates(head *listNode) *listNode {
	if head == nil {
		return head
	}

	var count = 0
	var left, right, behind = head, head.Next, head // behind 指向 right 的上一个节点
	for ; right != nil; right = right.Next {
		if left.Val == right.Val {
			count++
		} else {
			if count > 0 {
				if left == head {
					// 首部处理
					head = behind
				} else {
					left.Next = right
				}
				left = right
				count = 0
			} else {
				left = left.Next
			}
		}
		behind = right
	}
	if count > 0 {
		// 尾部处理
		if left == head {
			head = behind
		} else {
			left.Next = nil
		}
	}

	return head
}
