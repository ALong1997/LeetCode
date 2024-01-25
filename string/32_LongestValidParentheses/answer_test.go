package LeetCode

import (
	"fmt"
	"testing"
)

type question32 struct {
	para32
	ans32
}

type para32 struct {
	s string
}

type ans32 struct {
	ans int
}

func Test_Problem32(t *testing.T) {
	qs := []question32{{
		para32: para32{"(()"},
		ans32:  ans32{2},
	}, {
		para32: para32{")()())"},
		ans32:  ans32{4},
	}, {
		para32: para32{"()(()"},
		ans32:  ans32{2},
	}}

	fmt.Printf("------------------------LeetCode Problem 32------------------------\n")

	for _, q := range qs {
		a, p := q.ans32, q.para32
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, longestValidParentheses(p.s), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceLongestValidParentheses(p.s), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2LongestValidParentheses(p.s), a)
	}
	fmt.Printf("\n")
}
