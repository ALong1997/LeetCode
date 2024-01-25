package LeetCode

import (
	"fmt"
	"testing"
)

type question87 struct {
	para87
	ans87
}

type para87 struct {
	s1 string
	s2 string
}

type ans87 struct {
	ans bool
}

func Test_Problem87(t *testing.T) {
	qs := []question87{{
		para87: para87{"great", "rgeat"},
		ans87:  ans87{true},
	}, {
		para87: para87{"abcde", "caebd"},
		ans87:  ans87{false},
	}}

	fmt.Printf("------------------------LeetCode Problem 87------------------------\n")

	for _, q := range qs {
		a, p := q.ans87, q.para87
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isScramble(p.s1, p.s2), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, largestRectangleAreaHeight(p.heights), a)
	}
	fmt.Printf("\n")
}
