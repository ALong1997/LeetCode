package LeetCode

import (
	"fmt"
	"testing"
)

type question9 struct {
	para9
	ans9
}

type para9 struct {
	x int
}

type ans9 struct {
	ans bool
}

func Test_Problem9(t *testing.T) {
	qs := []question9{{
		para9: para9{121},
		ans9:  ans9{true},
	}, {
		para9: para9{-121},
		ans9:  ans9{false},
	}, {
		para9: para9{10},
		ans9:  ans9{false},
	}}

	fmt.Printf("------------------------LeetCode Problem 5------------------------\n")

	for _, q := range qs {
		a, p := q.ans9, q.para9
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isPalindrome(p.x), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceLongestPalindrome(p.s), a)
	}
	fmt.Printf("\n")
}
