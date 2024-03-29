package LeetCode

import "github.com/ALong1997/LeetCode/common"

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	head := &common.ListNode{Val: 0}
	n1, n2, carry, move := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		move.Next = &common.ListNode{Val: (n1 + n2 + carry) % 10}
		move = move.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
