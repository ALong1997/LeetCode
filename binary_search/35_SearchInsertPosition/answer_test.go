package LeetCode

import (
	"fmt"
	"testing"
)

type question35 struct {
	para35
	ans35
}

type para35 struct {
	nums   []int
	target int
}

type ans35 struct {
	ans int
}

func Test_Problem35(t *testing.T) {
	qs := []question35{{
		para35: para35{[]int{1, 3, 5, 6}, 5},
		ans35:  ans35{2},
	}, {
		para35: para35{[]int{1, 3, 5, 6}, 2},
		ans35:  ans35{1},
	}, {
		para35: para35{[]int{1, 3, 5, 6}, 7},
		ans35:  ans35{4},
	}}

	fmt.Printf("------------------------LeetCode Problem 35------------------------\n")

	for _, q := range qs {
		a, p := q.ans35, q.para35
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, searchInsert(p.nums, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceLongestValidParentheses(p.s), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2LongestValidParentheses(p.s), a)
	}
	fmt.Printf("\n")
}
