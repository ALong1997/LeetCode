package leetcode

import (
	"fmt"
	"testing"
)

type question86 struct {
	para86
	ans86
}

type para86 struct {
	head *listNode
	x    int
}

type ans86 struct {
	ans *listNode
}

func Test_Problem86(t *testing.T) {
	l1_6 := &listNode{2, nil}
	l1_5 := &listNode{5, l1_6}
	l1_4 := &listNode{2, l1_5}
	l1_3 := &listNode{3, l1_4}
	l1_2 := &listNode{4, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_6 := &listNode{5, nil}
	a1_5 := &listNode{3, a1_6}
	a1_4 := &listNode{4, a1_5}
	a1_3 := &listNode{2, a1_4}
	a1_2 := &listNode{2, a1_3}
	a1_1 := &listNode{1, a1_2}

	qs := []question86{{
		para86: para86{l1_1, 3},
		ans86:  ans86{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 86------------------------\n")

	for _, q := range qs {
		a, p := q.ans86, q.para86
		fmt.Print("【input】:")
		p.head.printlistNode()
		fmt.Print("       【output】:")
		partition(p.head, p.x).printlistNode()
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
