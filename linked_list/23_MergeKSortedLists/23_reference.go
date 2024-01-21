package leetcode

/*
解法: 考虑归并排序思想优化myAnswer，用分治的方法进行合并。

将 k 个链表配对并将同一对中的链表合并；
第一轮合并以后， k 个链表被合并成了 k/2​个链表，平均长度为 2n/k，然后是 k/4, k/8 个链表等等；
重复这一过程，直到我们得到了最终的有序链表。

时间复杂度：考虑递归「向上回升」的过程——第一轮合并 k/2 组链表，每一组的时间代价是 O(2n)；第二轮合并 k/4 组链表，每一组的时间代价是 O(4n)......
所以总的时间代价是 O(∑i=1,∞  k/(2^i) * n2^i)= O(kn×logk)，故渐进时间复杂度为 O(kn×logk)。
空间复杂度：递归会使用到 O(logk) 空间代价的栈空间。


结果: 执行用时 :4 ms 内存消耗 :5.3 MB
*/

func referenceMergeKLists(lists []*listNode) *listNode {
	return merge(lists, 0, len(lists)- 1)
}

func merge(lists []*listNode, l, r int) *listNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l+r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}