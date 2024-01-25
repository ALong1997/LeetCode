package LeetCode

import (
	"fmt"
	"testing"
)

type question58 struct {
	para58
	ans58
}

type para58 struct {
	s string
}

type ans58 struct {
	ans int
}

func Test_Problem58(t *testing.T) {
	qs := []question58{{
		para58: para58{"Hello World"},
		ans58:  ans58{5},
	}}

	fmt.Printf("------------------------LeetCode Problem 58------------------------\n")

	for _, q := range qs {
		a, p := q.ans58, q.para58
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, lengthOfLastWord(p.s), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
