package leetcode

import (
	"fmt"
	"testing"
)

type question59 struct {
	para59
	ans59
}

type para59 struct {
	n int
}

type ans59 struct {
	ans [][]int
}

func Test_Problem59(t *testing.T) {
	qs := []question59{{
		para59: para59{3},
		ans59:  ans59{[][]int{{1,2,3},{8,9,4},{7,6,5}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 59------------------------\n")

	for _, q := range qs {
		a, p := q.ans59, q.para59
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, generateMatrix(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
