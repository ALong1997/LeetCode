package leetcode

import (
	"fmt"
	"testing"
)

type question24 struct {
	para24
	ans24
}

type para24 struct {
	head *ListNode
}

type ans24 struct {
	ans *ListNode
}

func Test_Problem24(t *testing.T) {

	l1_4 := &ListNode{4, nil}
	l1_3 := &ListNode{3, l1_4}
	l1_2 := &ListNode{2, l1_3}
	l1_1 := &ListNode{1, l1_2}

	a1_4 := &ListNode{4, nil}
	a1_3 := &ListNode{3, a1_4}
	a1_2 := &ListNode{1, a1_3}
	a1_1 := &ListNode{2, a1_2}


	qs := []question24{{
		para24: para24{l1_1},
		ans24:  ans24{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 24------------------------\n")

	for _, q := range qs {
		a, p := q.ans24, q.para24
		fmt.Print("【input】:")
		p.head.printListNode()
		fmt.Print("       【output】:")
		swapPairs(p.head).printListNode()
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