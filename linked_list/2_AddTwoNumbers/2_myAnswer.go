package leetcode
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * Definition for singly-linked list.
 * type listNode struct {
 *     Val int
 *     Next *listNode
 * }
*/

type listNode struct {
	Val  int
	Next *listNode
}

/*
解法: 链表相当于把一个数的各位逆序处理，所以可以正序遍历并对应相加，当l1或l2其中一条链表遍历完，直接让l3.Next指向未结束的链表，注意进位。

结果: 执行用时 :4 ms 内存消耗 :5 MB
*/
func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	l3 := &listNode{0, nil}
	head := l3
	tail := l3
	carry, sum := 0, 0
	// l1 和 l2 非空
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		l3.Val = sum % 10
		carry = sum / 10
		l3.Next = &listNode{0, nil}
		tail = l3
		l1 = l1.Next
		l2 = l2.Next
		l3 = l3.Next
	}
	// l1 比 l2 长，l2 已遍历完
	if l1 != nil {
		tail.Next = l1
		l3 = l1
	}
	// l1 比 l2 短， l1 已遍历完
	if l2 != nil {
		tail.Next = l2
		l3 = l2
	}
	for ; l3.Next != nil; l3 = l3.Next {
		sum = l3.Val + carry
		l3.Val = sum % 10
		carry = sum / 10
		tail = l3
	}

	// 可能还有进位需要处理
	for ; carry != 0; l3 = l3.Next {
		sum = l3.Val + carry
		l3.Val = sum % 10
		carry = sum / 10
		if l3.Next == nil {
			l3.Next = &listNode{0, nil}
		}
		tail = l3
	}
	// 尾部处理
	if l3.Val == 0 && l3.Next == nil {
		tail.Next = nil
	}
	return head
}
