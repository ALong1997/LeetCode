package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	para26
	ans26
}

type para26 struct {
	nums []int
}

type ans26 struct {
	ans int
}

func Test_Problem26(t *testing.T) {
	qs := []question26{{
		para26: para26{[]int{1, 1, 2}},
		ans26:  ans26{2},
	}, {
		para26: para26{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
		ans26:  ans26{5},
	}}

	fmt.Printf("------------------------Leetcode Problem 13------------------------\n")

	for _, q := range qs {
		a, p := q.ans26, q.para26
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, removeDuplicates(p.nums), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}