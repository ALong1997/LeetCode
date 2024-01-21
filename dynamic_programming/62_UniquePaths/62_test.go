package leetcode

import (
	"fmt"
	"testing"
)

type question62 struct {
	para62
	ans62
}

type para62 struct {
	m int
	n int
}

type ans62 struct {
	ans int
}

func Test_Problem62(t *testing.T) {
	qs := []question62{{
		para62: para62{3, 2},
		ans62:  ans62{3},
	}, {
		para62: para62{7, 3},
		ans62:  ans62{28},
	}}

	fmt.Printf("------------------------Leetcode Problem 62------------------------\n")

	for _, q := range qs {
		a, p := q.ans62, q.para62
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, uniquePaths(p.m, p.n), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePaths(p.m, p.n), a)
	}
	fmt.Printf("\n")
}
