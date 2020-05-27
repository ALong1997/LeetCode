package leetcode

import (
	"fmt"
	"testing"
)

type question38 struct {
	para38
	ans38
}

type para38 struct {
	n int
}

type ans38 struct {
	ans string
}

func Test_Problem36(t *testing.T) {
	qs := []question38{{
		para38: para38{1},
		ans38:  ans38{"1"},
	}, {
		para38: para38{4},
		ans38:  ans38{"1211"},
	}}

	fmt.Printf("------------------------Leetcode Problem 38------------------------\n")

	for _, q := range qs {
		a, p := q.ans38, q.para38
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, countAndSay(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsValidSudoku(p.board), a)
	}
	fmt.Printf("\n")
}