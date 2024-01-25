package LeetCode

import (
	"fmt"
	"testing"
)

type question15 struct {
	para15
	ans15
}

type para15 struct {
	nums []int
}

type ans15 struct {
	ans [][]int
}

func Test_Problem15(t *testing.T) {
	qs := []question15{{
		para15: para15{[]int{-1, 0, 1, 2, -1, -4}},
		ans15:  ans15{[][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}}

	fmt.Printf("------------------------LeetCode Problem 15------------------------\n")

	for _, q := range qs {
		a, p := q.ans15, q.para15
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, threeSum(p.nums), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}
