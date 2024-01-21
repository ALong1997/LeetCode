package leetcode

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

type para46 struct {
	nums []int
}

type ans46 struct {
	ans [][]int
}

func Test_Problem46(t *testing.T) {
	qs := []question46{{
		para46: para46{[]int{1,2,3}},
		ans46:  ans46{[][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, q := range qs {
		a, p := q.ans46, q.para46
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, permute(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}