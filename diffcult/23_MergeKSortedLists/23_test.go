package leetcode

import (
	"fmt"
	"testing"
)

type question23 struct {
	para23
	ans23
}

type para23 struct {
	lists []*ListNode
}

type ans23 struct {
	ans *ListNode
}

func Test_Problem23(t *testing.T) {

	l1_3 := &ListNode{5, nil}
	l1_2 := &ListNode{4, l1_3}
	l1_1 := &ListNode{1, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{3, l2_3}
	l2_1 := &ListNode{1, l2_2}

	l3_2 := &ListNode{6, nil}
	l3_1 := &ListNode{2, l3_2}

	a1_8 := &ListNode{6, nil}
	a1_7 := &ListNode{5, a1_8}
	a1_6 := &ListNode{4, a1_7}
	a1_5 := &ListNode{4, a1_6}
	a1_4 := &ListNode{3, a1_5}
	a1_3 := &ListNode{2, a1_4}
	a1_2 := &ListNode{1, a1_3}
	a1_1 := &ListNode{1, a1_2}


	qs := []question23{{
		para23: para23{[]*ListNode{l1_1, l2_1, l3_1}},
		ans23:  ans23{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 23------------------------\n")

	for _, q := range qs {
		a, p := q.ans23, q.para23
		fmt.Print("【input】:")
		for _, e := range p.lists {
			e.printListNode()
		}
		fmt.Print("       【output】:")
		referenceMergeKLists(p.lists).printListNode()
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