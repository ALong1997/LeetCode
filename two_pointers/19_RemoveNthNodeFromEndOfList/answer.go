package LeetCode

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
你能尝试使用一趟扫描实现吗？
*/

/*
解法: 我们可以使用两个指针而不是一个指针。
第一个指针从列表的开头向前移动 n+1n+1 步，而第二个指针将从列表的开头出发。
现在，这两个指针被 nn 个结点分开。我们通过同时移动两个指针向前来保持这个恒定的间隔，直到第一个指针到达最后一个结点。
此时第二个指针将指向从最后一个结点数起的第 nn 个结点。我们重新链接第二个指针所引用的结点的 next 指针指向该结点的下下个结点。

结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

// Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

func removeNthFromEnd(head *listNode, n int) *listNode {
	var p, q *listNode
	var distance int
	q = head
	for q.Next != nil && distance < n {
		q = q.Next
		distance++
	}
	if distance < n-1 {
		return nil
	}
	if distance == n-1 {
		return head.Next
	}

	p = head
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return head
}
