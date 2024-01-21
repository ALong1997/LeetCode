package leetcode

import (
	"fmt"
	"testing"
)

type question52 struct {
	para52
	ans52
}

type para52 struct {
	n	int
}

type ans52 struct {
	ans int
}

func Test_Problem52(t *testing.T) {
	qs := []question52{{
		para52: para52{2},
		ans52:  ans52{2},
	}}

	fmt.Printf("------------------------Leetcode Problem 52------------------------\n")

	for _, q := range qs {
		a, p := q.ans52, q.para52
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, totalNQueens(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyPow(p.x, p.n), a)
	}
	fmt.Printf("\n")
}