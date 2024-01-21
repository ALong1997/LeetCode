package LeetCode

import (
	"fmt"
	"github.com/ALong1997/LeetCode/common"
	"testing"
)

type question2 struct {
	para2
	ans2
}

type para2 struct {
	l1 *common.ListNode
	l2 *common.ListNode
}

type ans2 struct {
	ans *common.ListNode
}

func Test_Problem2(t *testing.T) {
	l1_3 := &common.ListNode{Val: 3}
	l1_2 := &common.ListNode{Val: 4, Next: l1_3}
	l1_1 := &common.ListNode{Val: 2, Next: l1_2}

	l2_3 := &common.ListNode{Val: 4}
	l2_2 := &common.ListNode{Val: 6, Next: l2_3}
	l2_1 := &common.ListNode{Val: 5, Next: l2_2}

	a12_3 := &common.ListNode{Val: 8}
	a12_2 := &common.ListNode{Next: a12_3}
	a12_1 := &common.ListNode{Val: 7, Next: a12_2}

	l3_1 := &common.ListNode{Val: 5}

	l4_1 := &common.ListNode{Val: 5}

	a34_2 := &common.ListNode{Val: 1}
	a34_1 := &common.ListNode{Next: a34_2}

	l5_4 := &common.ListNode{Val: 1}
	l5_3 := &common.ListNode{Next: l5_4}
	l5_2 := &common.ListNode{Val: 8, Next: l5_3}
	l5_1 := &common.ListNode{Val: 8, Next: l5_2}

	l6_1 := &common.ListNode{Val: 2}

	a56_4 := &common.ListNode{Val: 1}
	a56_3 := &common.ListNode{Next: a56_4}
	a56_2 := &common.ListNode{Val: 9, Next: a56_3}
	a56_1 := &common.ListNode{Next: a56_2}

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
		p.l1.Print()
		p.l2.Print()
		fmt.Print("       【output】:")
		addTwoNumbers(p.l1, p.l2).Print()
		fmt.Print("       【answer】:")
		a.ans.Print()
		fmt.Printf("\n")
		//fmt.Printf("【input】:%p       【output】:%p       【answer】:%p\n", p, addTwoNumbers(p.l1, p.l2), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceTwoSum(p.nums, p.target), a)
	}
	fmt.Printf("\n")
}
