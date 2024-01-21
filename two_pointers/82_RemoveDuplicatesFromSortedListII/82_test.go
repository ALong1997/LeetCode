package LeetCode

import (
	"fmt"
	"testing"
)

type question82 struct {
	para82
	ans82
}

type para82 struct {
	head *listNode
}

type ans82 struct {
	ans *listNode
}

func Test_Problem82(t *testing.T) {
	l1_7 := &listNode{5, nil}
	l1_6 := &listNode{4, l1_7}
	l1_5 := &listNode{4, l1_6}
	l1_4 := &listNode{3, l1_5}
	l1_3 := &listNode{3, l1_4}
	l1_2 := &listNode{2, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_3 := &listNode{5, nil}
	a1_2 := &listNode{2, a1_3}
	a1_1 := &listNode{1, a1_2}

	l2_5 := &listNode{3, nil}
	l2_4 := &listNode{2, l2_5}
	l2_3 := &listNode{1, l2_4}
	l2_2 := &listNode{1, l2_3}
	l2_1 := &listNode{1, l2_2}

	a2_2 := &listNode{3, nil}
	a2_1 := &listNode{2, a2_2}

	qs := []question82{{
		para82: para82{l1_1},
		ans82:  ans82{a1_1},
	}, {
		para82: para82{l2_1},
		ans82:  ans82{a2_1},
	}}

	fmt.Printf("------------------------LeetCode Problem 82------------------------\n")

	for _, q := range qs {
		a, p := q.ans82, q.para82
		fmt.Print("【input】:")
		p.head.printlistNode()
		fmt.Print("       【output】:")
		deleteDuplicates(p.head).printlistNode()
		fmt.Print("       【answer】:")
		a.ans.printlistNode()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, threeSum(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}

// 输出 listNode ：(2 -> 4 -> 3)
func (l *listNode) printlistNode() {
	fmt.Print("(")
	for ; l != nil; l = l.Next {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Print(")")
}
