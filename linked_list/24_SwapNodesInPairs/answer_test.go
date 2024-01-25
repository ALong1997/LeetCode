package LeetCode

import (
	"fmt"
	"testing"
)

type question24 struct {
	para24
	ans24
}

type para24 struct {
	head *listNode
}

type ans24 struct {
	ans *listNode
}

func Test_Problem24(t *testing.T) {

	l1_4 := &listNode{4, nil}
	l1_3 := &listNode{3, l1_4}
	l1_2 := &listNode{2, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_4 := &listNode{3, nil}
	a1_3 := &listNode{4, a1_4}
	a1_2 := &listNode{1, a1_3}
	a1_1 := &listNode{2, a1_2}

	qs := []question24{{
		para24: para24{l1_1},
		ans24:  ans24{a1_1},
	}}

	fmt.Printf("------------------------LeetCode Problem 24------------------------\n")

	for _, q := range qs {
		a, p := q.ans24, q.para24
		fmt.Print("【input】:")
		p.head.printlistNode()
		fmt.Print("       【output】:")
		swapPairs(p.head).printlistNode()
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
