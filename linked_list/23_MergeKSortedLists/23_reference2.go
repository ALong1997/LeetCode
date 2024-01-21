package LeetCode

import "container/heap"

/*
解法: 最小堆
由于每个链表都是递增的，因此每次要拿的最小值都是在每个链表的头节点。
比如，上面3个有序链表第一个值应该在三个链表的头节点1，1，0中找最小值为0。
之后游标移动到 0 -> 2 节点0之后的节点2，然后要找1，1，2中的最小值。
使用这个方法依次取每个链表上的每一个节点，直到取完，就可以得到合并后的有序链表。

这种方法每一轮要对比k次。每一轮对比完可以确定一个节点，总共有n个节点。
因此时间复杂度是O(k*n)，这里可以做一点优化，使用最小堆来维护当前k个值的关系。
这样每次只需要O(1)的时间取得当前k个值的最小值。

则新的节点加入到大小为k的堆，则需要O(logK)的时间去调整堆的结构。
这样每一轮操作只能确定一个节点的位置，因此总时间复杂度是O(n*logk),
使用了大小为k的辅助堆，因此空间复杂度是O(k)


结果: 执行用时 :12 ms 内存消耗 :5.9 MB
*/

// Time: O(n*log(k)), Space: O(k) n是总节点数量，k是链表个数
func reference2mergeKLists(lists []*listNode) *listNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var h intHeap // 最小堆用于维护当前k个节点
	heap.Init(&h) // 用于节点间的比较

	for _, list := range lists {
		// 数组中非空的链表加入到最小堆
		if list != nil {
			heap.Push(&h, list)
		}
	}
	// 定义dummy节点用于统一处理
	dummy := &listNode{}
	p := dummy // p初始指向dummy节点

	// 当最小堆不为空时，不断执行以下操作
	for h.Len() > 0 { // 取出堆顶元素，即取出最小值节点
		min := heap.Pop(&h).(*listNode)
		p.Next = min // 游标p指向最小值节点
		p = p.Next   // p向后移动一个位置
		// 这样就确定一个节点在合并链表中的位置
		if min.Next != nil { // 如果最小值节点后面的节点非空
			heap.Push(&h, min.Next) // 则把最小值节点后面的节点加入到最小堆中
		}
	}
	return dummy.Next // 最后只要返回dummy.Next即可
}

type intHeap []*listNode

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(*listNode)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
