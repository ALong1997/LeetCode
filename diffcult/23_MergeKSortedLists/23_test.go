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
	lists []*listNode
}

type ans23 struct {
	ans *listNode
}

func Test_Problem23(t *testing.T) {

	l1_3 := &listNode{5, nil}
	l1_2 := &listNode{4, l1_3}
	l1_1 := &listNode{1, l1_2}

	l2_3 := &listNode{4, nil}
	l2_2 := &listNode{3, l2_3}
	l2_1 := &listNode{1, l2_2}

	l3_2 := &listNode{6, nil}
	l3_1 := &listNode{2, l3_2}

	a1_8 := &listNode{6, nil}
	a1_7 := &listNode{5, a1_8}
	a1_6 := &listNode{4, a1_7}
	a1_5 := &listNode{4, a1_6}
	a1_4 := &listNode{3, a1_5}
	a1_3 := &listNode{2, a1_4}
	a1_2 := &listNode{1, a1_3}
	a1_1 := &listNode{1, a1_2}


	qs := []question23{{
		para23: para23{[]*listNode{l1_1, l2_1, l3_1}},
		ans23:  ans23{a1_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 23------------------------\n")

	for _, q := range qs {
		a, p := q.ans23, q.para23
		fmt.Print("【input】:")
		for _, e := range p.lists {
			e.printlistNode()
		}
		fmt.Print("       【output】:")
		reference2mergeKLists(p.lists).printlistNode()
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