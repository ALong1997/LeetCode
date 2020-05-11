package leetcode

import (
	"fmt"
	"testing"
)

type question22 struct {
	para22
	ans22
}

type para22 struct {
	n int
}

type ans22 struct {
	ans []string
}

func Test_Problem22(t *testing.T) {
	qs := []question22{{
		para22: para22{3},
		ans22:  ans22{[]string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}}

	fmt.Printf("------------------------Leetcode Problem 22------------------------\n")

	for _, q := range qs {
		a, p := q.ans22, q.para22
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, generateParenthesis(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}