package LeetCode

import (
	"fmt"
	"testing"
)

type question45 struct {
	para45
	ans45
}

type para45 struct {
	nums []int
}

type ans45 struct {
	ans int
}

func Test_Problem45(t *testing.T) {
	qs := []question45{{
		para45: para45{[]int{2, 3, 1, 1, 4}},
		ans45:  ans45{2},
	}, {
		para45: para45{[]int{1, 2, 1, 1, 1}},
		ans45:  ans45{3},
	}}

	fmt.Printf("------------------------LeetCode Problem 45------------------------\n")

	for _, q := range qs {
		a, p := q.ans45, q.para45
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, jump(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}
