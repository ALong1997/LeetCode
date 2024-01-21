package LeetCode

import (
	"fmt"
	"testing"
)

type question55 struct {
	para55
	ans55
}

type para55 struct {
	nums []int
}

type ans55 struct {
	ans bool
}

func Test_Problem55(t *testing.T) {
	qs := []question55{{
		para55: para55{[]int{2, 3, 1, 1, 4}},
		ans55:  ans55{true},
	}, {
		para55: para55{[]int{3, 2, 1, 0, 4}},
		ans55:  ans55{false},
	}, {
		para55: para55{[]int{2, 5, 0, 0}},
		ans55:  ans55{true},
	}}

	fmt.Printf("------------------------LeetCode Problem 55------------------------\n")

	for _, q := range qs {
		a, p := q.ans55, q.para55
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, canJump(p.nums), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
