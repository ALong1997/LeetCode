/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
package leetcode

/*
解法: 暴力解


结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var ans, p, q, r *ListNode
	r = &ListNode{
		Val:  0,
		Next: head,
	}
	ans = r
	for p = head ; p != nil; p = p.Next {
		q = p.Next
		if q == nil {
			break
		}
		p.Next = q.Next
		q.Next = p
		r.Next = q
		r = p
	}
	return ans.Next
}