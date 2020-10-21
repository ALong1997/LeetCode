/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
*/
package leetcode

/*
解法: 双指针


结果: 执行用时 :4 ms 内存消耗 :3 MB
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
	var p, q, r = head, head.Next, head
	for ; q != nil; q = q.Next {
		if p.Val == q.Val {
			count++
		} else {
			if count > 0 {
				if p == head {
					// 首部处理
					head = q
				} else {
					r.Next = q
				}
				p = q
				count = 0
			} else {
				r = p
				p = p.Next
			}
		}
	}
	if count > 0 {
		// 尾部处理
		if p == head {
			head = q
		} else {
			r.Next = q
		}
	}

	return head
}
