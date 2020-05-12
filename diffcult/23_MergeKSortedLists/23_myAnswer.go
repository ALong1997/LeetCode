/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
*/
package leetcode

/*
解法: 利用T21合并两个有序链表的答案

时间复杂度：假设每个链表的最长长度是 n。在第一次合并后，ans 的长度为 n；第二次合并后，ans 的长度为 2n，第 i 次合并后，ans 的长度为 in。
第 i 次合并的时间代价是 O(n+(i−1)×n)=O(i×n)，那么总的时间代价为 O(∑i=1,k(i×n)) = O(nk^2)
空间复杂度：没有用到与 k 和 n 规模相关的辅助空间，故渐进空间复杂度为 O(1)。


结果: 执行用时 :128 ms 内存消耗 :5.3 MB
*/

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists) ; i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
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