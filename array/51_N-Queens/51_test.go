package leetcode

import (
	"fmt"
	"testing"
)

type question51 struct {
	para51
	ans51
}

type para51 struct {
	n	int
}

type ans51 struct {
	ans [][]string
}

func Test_Problem51(t *testing.T) {
	qs := []question51{{
		para51: para51{4},
		ans51:  ans51{[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 51------------------------\n")

	for _, q := range qs {
		a, p := q.ans51, q.para51
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, solveNQueens(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyPow(p.x, p.n), a)
	}
	fmt.Printf("\n")
}