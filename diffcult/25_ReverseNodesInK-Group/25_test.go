package leetcode

import (
	"fmt"
	"testing"
)

type question25 struct {
	para25
	ans25
}

type para25 struct {
	head *listNode
	k int
}

type ans25 struct {
	ans *listNode
}

func Test_Problem25(t *testing.T) {

	l1_5 := &listNode{5, nil}
	l1_4 := &listNode{4, l1_5}
	l1_3 := &listNode{3, l1_4}
	l1_2 := &listNode{2, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_5 := &listNode{5, nil}
	a1_4 := &listNode{3, a1_5}
	a1_3 := &listNode{4, a1_4}
	a1_2 := &listNode{1, a1_3}
	a1_1 := &listNode{2, a1_2}


	qs := []question25{{
		para25: para25{l1_1, 2},
		ans25:  ans25{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 24------------------------\n")

	for _, q := range qs {
		a, p := q.ans25, q.para25
		fmt.Print("【input】:")
		p.head.printlistNode()
		fmt.Print("       【output】:")
		reverseKGroup(p.head, p.k).printlistNode()
		fmt.Print("       【answer】:")
		a.ans.printlistNode()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, threeSum(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}

// 输出 listNode ：(2 -> 4 -> 3)
func (l *listNode) printlistNode()  {
	fmt.Print("(")
	for ; l != nil ; l = l.Next {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Print(")")
}