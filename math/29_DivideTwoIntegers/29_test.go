package leetcode

import (
	"fmt"
	"testing"
)

type question29 struct {
	para29
	ans29
}

type para29 struct {
	dividend int
	divisor int
}

type ans29 struct {
	ans int
}

func Test_Problem29(t *testing.T) {
	qs := []question29{{
		para29: para29{10, 3},
		ans29:  ans29{3},
	}, {
		para29: para29{7, -3},
		ans29:  ans29{-2},
	}}

	fmt.Printf("------------------------Leetcode Problem 29------------------------\n")

	for _, q := range qs {
		a, p := q.ans29, q.para29
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, divide(p.dividend, p.divisor), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceStrStr(p.aystack, p.needle), a)
	}
	fmt.Printf("\n")
}