package leetcode

import (
	"fmt"
	"testing"
)

type question81 struct {
	para81
	ans81
}

type para81 struct {
	nums   []int
	target int
}

type ans81 struct {
	ans bool
}

func Test_Problem81(t *testing.T) {
	qs := []question81{{
		para81: para81{[]int{2, 5, 6, 0, 0, 1, 2}, 0},
		ans81:  ans81{true},
	}, {
		para81: para81{[]int{2, 5, 6, 0, 0, 1, 2}, 3},
		ans81:  ans81{false},
	}, {
		para81: para81{[]int{1, 3, 1, 1, 1, 1}, 3},
		ans81:  ans81{true},
	}, {
		para81: para81{[]int{3, 1}, 1},
		ans81:  ans81{true},
	}}

	fmt.Printf("------------------------Leetcode Problem 81------------------------\n")

	for _, q := range qs {
		a, p := q.ans81, q.para81
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, search(p.nums, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceSubsets(p.nums), a)
	}
	fmt.Printf("\n")
}
