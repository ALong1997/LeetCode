package leetcode

import (
	"fmt"
	"testing"
)

type question60 struct {
	para60
	ans60
}

type para60 struct {
	n int
	k int
}

type ans60 struct {
	ans string
}

func Test_Problem60(t *testing.T) {
	qs := []question60{{
		para60: para60{3, 3},
		ans60:  ans60{"213"},
	}, {
		para60: para60{4, 9},
		ans60:  ans60{"2314"},
	}, {
		para60: para60{3, 2},
		ans60:  ans60{"132"},
	}}

	fmt.Printf("------------------------Leetcode Problem 60------------------------\n")

	for _, q := range qs {
		a, p := q.ans60, q.para60
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, getPermutation(p.n, p.k), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
