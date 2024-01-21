package LeetCode

import (
	"fmt"
	"testing"
)

type question18 struct {
	para18
	ans18
}

type para18 struct {
	nums   []int
	target int
}

type ans18 struct {
	ans [][]int
}

func Test_Problem18(t *testing.T) {
	qs := []question18{{
		para18: para18{[]int{1, 0, -1, 0, -2, 2}, 0},
		ans18:  ans18{[][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
	}, {
		para18: para18{[]int{-3, -1, 0, 2, 4, 5}, 2},
		ans18:  ans18{[][]int{{-3, -1, 2, 4}}},
	}}

	fmt.Printf("------------------------LeetCode Problem 18------------------------\n")

	for _, q := range qs {
		a, p := q.ans18, q.para18
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, fourSum(p.nums, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}
