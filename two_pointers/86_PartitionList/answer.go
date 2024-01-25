package LeetCode

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
*/

/*
解法: 暴力解
先找到第一个大于或等于 x 的值，记录其位置，紧接着，将小于 x 的挪到该位置左边即可。


结果: 执行用时 :0 ms 内存消耗 :2.4 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

func partition(head *listNode, x int) *listNode {
	var p = &listNode{Next: head} // 头节点
	var q, divide, before = head, head, p
	// 先找到第一个大于或等于 x 的值
	for q != nil {
		if q.Val >= x {
			break
		}
		q, divide, before = q.Next, divide.Next, before.Next
	}
	// 将小于 x 的挪到该位置左边即可
	if q != nil {
		var r = q.Next
		for r != nil {
			if r.Val < x {
				q.Next = r.Next
				before.Next = r
				r.Next = divide
				before = r
				r = q.Next
			} else {
				q, r = q.Next, r.Next
			}
		}
	}
	return p.Next
}
