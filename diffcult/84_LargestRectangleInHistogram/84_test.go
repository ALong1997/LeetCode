package leetcode

import (
	"fmt"
	"testing"
)

type question84 struct {
	para84
	ans84
}

type para84 struct {
	heights []int
}

type ans84 struct {
	ans int
}

func Test_Problem84(t *testing.T) {
	qs := []question84{{
		para84: para84{[]int{2,1,5,6,2,3}},
		ans84:  ans84{10},
	}}

	fmt.Printf("------------------------Leetcode Problem 84------------------------\n")

	for _, q := range qs {
		a, p := q.ans84, q.para84
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, largestRectangleArea(p.heights), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
