/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package leetcode

/*
解法: 暴力解


结果: 执行用时 :4 ms 内存消耗 :2.5 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val int
	Next *listNode
}

func mergeTwoLists(l1 *listNode, l2 *listNode) *listNode {
	dummy := &listNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
}