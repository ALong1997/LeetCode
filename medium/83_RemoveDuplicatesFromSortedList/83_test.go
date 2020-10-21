package leetcode

import (
	"fmt"
	"testing"
)

type question83 struct {
	para83
	ans83
}

type para83 struct {
	head *listNode
}

type ans83 struct {
	ans *listNode
}

func Test_Problem83(t *testing.T) {
	l1_3 := &listNode{2, nil}
	l1_2 := &listNode{1, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_2 := &listNode{2, nil}
	a1_1 := &listNode{1, a1_2}

	l2_5 := &listNode{3, nil}
	l2_4 := &listNode{3, l2_5}
	l2_3 := &listNode{2, l2_4}
	l2_2 := &listNode{1, l2_3}
	l2_1 := &listNode{1, l2_2}

	a2_3 := &listNode{3, nil}
	a2_2 := &listNode{2, a2_3}
	a2_1 := &listNode{1, a2_2}

	qs := []question83{{
		para83: para83{l1_1},
		ans83:  ans83{a1_1},
	}, {
		para83: para83{l2_1},
		ans83:  ans83{a2_1},
	}}

	fmt.Printf("------------------------Leetcode Problem 83------------------------\n")

	for _, q := range qs {
		a, p := q.ans83, q.para83
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
