package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Print (2 -> 4 -> 3)
func (l *ListNode) Print() {
	head := l
	fmt.Print("(")
	for ; head != nil; head = head.Next {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Print(")")
}
