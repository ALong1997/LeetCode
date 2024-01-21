package LeetCode

import (
	"fmt"
	"testing"
)

type question2 struct {
	para2
	ans2
}

type para2 struct {
	l1 *listNode
	l2 *listNode
}

type ans2 struct {
	ans *listNode
}

func Test_Problem2(t *testing.T) {
	l1_3 := &listNode{3, nil}
	l1_2 := &listNode{4, l1_3}
	l1_1 := &listNode{2, l1_2}

	l2_3 := &listNode{4, nil}
	l2_2 := &listNode{6, l2_3}
	l2_1 := &listNode{5, l2_2}

	a12_3 := &listNode{8, nil}
	a12_2 := &listNode{0, a12_3}
	a12_1 := &listNode{7, a12_2}

	l3_1 := &listNode{5, nil}

	l4_1 := &listNode{5, nil}

	a34_2 := &listNode{1, nil}
	a34_1 := &listNode{0, a34_2}

	l5_4 := &listNode{1, nil}
	l5_3 := &listNode{0, l5_4}
	l5_2 := &listNode{8, l5_3}
	l5_1 := &listNode{8, l5_2}

	l6_1 := &listNode{2, nil}

	a56_4 := &listNode{1, nil}
	a56_3 := &listNode{0, a56_4}
	a56_2 := &listNode{9, a56_3}
	a56_1 := &listNode{0, a56_2}

	qs := []question2{{
		para2: para2{l1_1, l2_1},
		ans2:  ans2{a12_1},
	}, {
		para2: para2{l3_1, l4_1},
		ans2:  ans2{a34_1},
	}, {
		para2: para2{l5_1, l6_1},
		ans2:  ans2{a56_1},
	}}

	fmt.Printf("------------------------LeetCode Problem 2------------------------\n")

	for _, q := range qs {
		a, p := q.ans2, q.para2
		fmt.Print("【input】:")
		p.l1.printlistNode()
		p.l2.printlistNode()
		fmt.Print("       【output】:")
		addTwoNumbers(p.l1, p.l2).printlistNode()
		fmt.Print("       【answer】:")
		a.ans.printlistNode()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%p       【output】:%p       【answer】:%p\n", p, addTwoNumbers(p.l1, p.l2), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceTwoSum(p.nums, p.target), a)
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
