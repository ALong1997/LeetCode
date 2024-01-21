package LeetCode

import (
	"fmt"
	"testing"
)

type question69 struct {
	para69
	ans69
}

type para69 struct {
	x int
}

type ans69 struct {
	ans int
}

func Test_Problem69(t *testing.T) {
	qs := []question69{{
		para69: para69{4},
		ans69:  ans69{2},
	}, {
		para69: para69{8},
		ans69:  ans69{2},
	}}

	fmt.Printf("------------------------LeetCode Problem 69------------------------\n")

	for _, q := range qs {
		a, p := q.ans69, q.para69
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, mySqrt(p.x), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
