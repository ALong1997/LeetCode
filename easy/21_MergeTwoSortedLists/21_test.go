package leetcode

import (
	"fmt"
	"testing"
)

type question21 struct {
	para21
	ans21
}

type para21 struct {
	l1 *ListNode
	l2 *ListNode
}

type ans21 struct {
	ans *ListNode
}

func Test_Problem21(t *testing.T) {

	l1_3 := &ListNode{4, nil}
	l1_2 := &ListNode{2, l1_3}
	l1_1 := &ListNode{1, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{3, l2_3}
	l2_1 := &ListNode{1, l2_2}

	a1_6 := &ListNode{4, nil}
	a1_5 := &ListNode{4, a1_6}
	a1_4 := &ListNode{3, a1_5}
	a1_3 := &ListNode{2, a1_4}
	a1_2 := &ListNode{1, a1_3}
	a1_1 := &ListNode{1, a1_2}


	qs := []question21{{
		para21: para21{l1_1, l2_1},
		ans21:  ans21{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		a, p := q.ans21, q.para21
		fmt.Print("【input】:")
		p.l1.printListNode()
		p.l2.printListNode()
		fmt.Print("       【output】:")
		mergeTwoLists(p.l1, p.l2).printListNode()
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