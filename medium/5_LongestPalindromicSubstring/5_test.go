package leetcode

import (
	"fmt"
	"testing"
)

type question5 struct {
	para5
	ans5
}

type para5 struct {
	s	string
}

type ans5 struct {
	s []string
}

func Test_Problem5(t *testing.T) {
	qs := []question5{{
		para5: para5{"babad"},
		ans5:  ans5{[]string{"bab", "aba"}},
	}, {
		para5: para5{"cbbd"},
		ans5:  ans5{[]string{"bb"}},
	}}

	fmt.Printf("------------------------Leetcode Problem 5------------------------\n")

	for _, q := range qs {
		a, p := q.ans5, q.para5
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, longestPalindrome(p.s), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceFindMedianSortedArrays(p.nums1, p.nums2), a)
	}
	fmt.Printf("\n")
}