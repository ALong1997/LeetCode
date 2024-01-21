package leetcode

import (
	"fmt"
	"testing"
)

type question78 struct {
	para78
	ans78
}

type para78 struct {
	nums []int
}

type ans78 struct {
	ans [][]int
}

func Test_Problem78(t *testing.T) {
	qs := []question78{{
		para78: para78{[]int{1, 2, 3}},
		ans78:  ans78{[][]int{{}, {1}, {2}, {3}, {2, 3}, {1, 2}, {1, 3}, {1, 2, 3}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 78------------------------\n")

	for _, q := range qs {
		a, p := q.ans78, q.para78
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, subsets(p.nums), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceSubsets(p.nums), a)
	}
	fmt.Printf("\n")
}
