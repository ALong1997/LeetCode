package leetcode
/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
*/

/*
解法: 将第 k 个位置后的节点放到头部即可


结果: 执行用时 :4 ms 内存消耗 :2.5 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

func rotateRight(head *listNode, k int) *listNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 先变成首尾相连的环状链表
	length, cycle := 1, head
	for cycle.Next != nil {
		length++
		cycle = cycle.Next
	}
	cycle.Next = head

	// 找到倒数第k+1个结点
	p := head
	for i := 0; i < length-k%length-1; i++ {
		p = p.Next
	}

	// 分割
	ans := p.Next
	p.Next = nil
	return ans
}
