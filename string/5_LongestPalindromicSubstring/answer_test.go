package LeetCode

import (
	"fmt"
	"testing"
)

type question5 struct {
	para5
	ans5
}

type para5 struct {
	s string
}

type ans5 struct {
	ans []string
}

func Test_Problem5(t *testing.T) {
	qs := []question5{{
		para5: para5{"babad"},
		ans5:  ans5{[]string{"bab", "aba"}},
	}, {
		para5: para5{"cbbd"},
		ans5:  ans5{[]string{"bb"}},
	}}

	fmt.Printf("------------------------LeetCode Problem 5------------------------\n")

	for _, q := range qs {
		a, p := q.ans5, q.para5
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, longestPalindrome(p.s), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, LongestPalindrome(p.s), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2LongestPalindrome(p.s), a)
	}
	fmt.Printf("\n")
}
