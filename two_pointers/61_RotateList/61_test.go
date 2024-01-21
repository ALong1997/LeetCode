package LeetCode

import (
	"fmt"
	"testing"
)

type question61 struct {
	para61
	ans61
}

type para61 struct {
	head *listNode
	k    int
}

type ans61 struct {
	ans *listNode
}

func Test_Problem61(t *testing.T) {

	l1_5 := &listNode{5, nil}
	l1_4 := &listNode{4, l1_5}
	l1_3 := &listNode{3, l1_4}
	l1_2 := &listNode{2, l1_3}
	l1_1 := &listNode{1, l1_2}

	a1_5 := &listNode{3, nil}
	a1_4 := &listNode{2, a1_5}
	a1_3 := &listNode{1, a1_4}
	a1_2 := &listNode{5, a1_3}
	a1_1 := &listNode{4, a1_2}

	l2_3 := &listNode{2, nil}
	l2_2 := &listNode{1, l2_3}
	l2_1 := &listNode{0, l2_2}

	a2_3 := &listNode{1, nil}
	a2_2 := &listNode{0, a2_3}
	a2_1 := &listNode{2, a2_2}

	qs := []question61{{
		para61: para61{l1_1, 2},
		ans61:  ans61{a1_1},
	}, {
		para61: para61{l2_1, 4},
		ans61:  ans61{a2_1},
	}}

	fmt.Printf("------------------------LeetCode Problem 61------------------------\n")

	for _, q := range qs {
		a, p := q.ans61, q.para61
		fmt.Print("【input】:")
		p.head.printlistNode()
		fmt.Print(p.k, "       【output】:")
		rotateRight(p.head, p.k).printlistNode()
		fmt.Print("       【answer】:")
		a.ans.printlistNode()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, rotateRight(p.head, p.k), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
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
