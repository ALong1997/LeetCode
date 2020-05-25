package leetcode

import (
	"fmt"
	"testing"
)

type question34 struct {
	para34
	ans34
}

type para34 struct {
	nums []int
	target int
}

type ans34 struct {
	ans []int
}

func Test_Problem34(t *testing.T) {
	qs := []question34{{
		para34: para34{[]int{5,7,7,8,8,10}, 8},
		ans34:  ans34{[]int{3, 4}},
	}, {
		para34: para34{[]int{5,7,7,8,8,10}, 6},
		ans34:  ans34{[]int{-1, -1}},
	}}

	fmt.Printf("------------------------Leetcode Problem 34------------------------\n")

	for _, q := range qs {
		a, p := q.ans34, q.para34
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, searchRange(p.nums,p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceLongestValidParentheses(p.s), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2LongestValidParentheses(p.s), a)
	}
	fmt.Printf("\n")
}