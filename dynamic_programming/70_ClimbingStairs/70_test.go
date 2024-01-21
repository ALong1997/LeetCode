package LeetCode

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

type para70 struct {
	n int
}

type ans70 struct {
	ans int
}

func Test_Problem70(t *testing.T) {
	qs := []question70{{
		para70: para70{2},
		ans70:  ans70{2},
	}, {
		para70: para70{3},
		ans70:  ans70{3},
	}}

	fmt.Printf("------------------------LeetCode Problem 70------------------------\n")

	for _, q := range qs {
		a, p := q.ans70, q.para70
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, climbStairs(p.n), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
