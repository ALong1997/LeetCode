package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

type para7 struct {
	x	int
}

type ans7 struct {
	ans int
}

func Test_Problem6(t *testing.T) {
	qs := []question7{{
		para7: para7{123},
		ans7:  ans7{321},
	}, {
		para7: para7{-123},
		ans7:  ans7{-321},
	}, {
		para7: para7{120},
		ans7:  ans7{21},
	}}

	fmt.Printf("------------------------Leetcode Problem 6------------------------\n")

	for _, q := range qs {
		a, p := q.ans7, q.para7
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reverse(p.x), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceConvert(p.s, p.numRows), a)
	}
	fmt.Printf("\n")
}