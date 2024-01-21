package leetcode
/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/*
解法: 暴力解, T24的拓展


结果: 执行用时 :4 ms 内存消耗 :3.6 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val int
	Next *listNode
}

func reverseKGroup(head *listNode, k int) *listNode {
	dummy := &listNode{
		Val: -1,
		Next: head,
	}
	pre := dummy
	cur := dummy.Next

	for {
		n := k
		// 找出下个部分的头
		nextPart := cur
		for nextPart != nil && n > 0{
			nextPart = nextPart.Next
			n--
		}
		// 如果已经不够 k 个说明可以返回了
		if n > 0{
			break
		} else {
			n = k
		}
		// 保存下个 Pre 节点
		nextPre := cur

		for n > 0{
			// 保存下当前元素的下一个元素
			temp := cur.Next
			// 接上下个头
			cur.Next = nextPart
			// 下个头为当前元素
			nextPart = cur
			// 当前元素为之前临时保留的元素
			cur = temp
			n--
		}
		// n次翻转完毕
		pre.Next = nextPart
		// 设置下个Pre
		pre = nextPre
		cur = pre.Next
	}
	return dummy.Next
}