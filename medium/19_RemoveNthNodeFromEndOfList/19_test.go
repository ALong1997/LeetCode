package leetcode

import (
	"fmt"
	"testing"
)

type question19 struct {
	para19
	ans19
}

type para19 struct {
	head *ListNode
	n int
}

type ans19 struct {
	ans *ListNode
}

func Test_Problem19(t *testing.T) {

	l1_5 := &ListNode{5, nil}
	l1_4 := &ListNode{4, l1_5}
	l1_3 := &ListNode{3, l1_4}
	l1_2 := &ListNode{2, l1_3}
	l1_1 := &ListNode{1, l1_2}

	a1_4 := &ListNode{5, nil}
	a1_3 := &ListNode{3, a1_4}
	a1_2 := &ListNode{2, a1_3}
	a1_1 := &ListNode{1, a1_2}

	l2_2 := &ListNode{2, nil}
	l2_1 := &ListNode{1, l2_2}

	a2_1 := &ListNode{1, nil}

	qs := []question19{{
		para19: para19{l1_1, 2},
		ans19:  ans19{a1_1},
	},{
		para19: para19{l2_1, 1},
		ans19:  ans19{a2_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 15------------------------\n")

	for _, q := range qs {
		a, p := q.ans19, q.para19
		fmt.Print("【input】:")
		p.head.printListNode()
		fmt.Print("       【output】:")
		removeNthFromEnd(p.head, p.n).printListNode()
		fmt.Print("       【answer】:")
		a.ans.printListNode()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, threeSum(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}

// 输出 ListNode ：(2 -> 4 -> 3)
func (l *ListNode) printListNode()  {
	fmt.Print("(")
	for ; l != nil ; l = l.Next {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Print(")")
}