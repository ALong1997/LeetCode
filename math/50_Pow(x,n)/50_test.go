package leetcode

import (
	"fmt"
	"testing"
)

type question50 struct {
	para50
	ans50
}

type para50 struct {
	x	float64
	n	int
}

type ans50 struct {
	ans float64
}

func Test_Problem50(t *testing.T) {
	qs := []question50{{
		para50: para50{2.00000, 10},
		ans50:  ans50{1024.00000},
	}, {
		para50: para50{2.10000, 3},
		ans50:  ans50{9.26100},
	}, {
		para50: para50{2.00000, -2},
		ans50:  ans50{0.25000},
	}}

	fmt.Printf("------------------------Leetcode Problem 50------------------------\n")

	for _, q := range qs {
		a, p := q.ans50, q.para50
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, myPow(p.x, p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyPow(p.x, p.n), a)
	}
	fmt.Printf("\n")
}