package LeetCode

import (
	"fmt"
	"testing"
)

type question20 struct {
	para20
	ans20
}

type para20 struct {
	s string
}

type ans20 struct {
	ans bool
}

func Test_Problem20(t *testing.T) {
	qs := []question20{{
		para20: para20{"()"},
		ans20:  ans20{true},
	}, {
		para20: para20{"()[]{}"},
		ans20:  ans20{true},
	}, {
		para20: para20{"(]"},
		ans20:  ans20{false},
	}, {
		para20: para20{"([)]"},
		ans20:  ans20{false},
	}}

	fmt.Printf("------------------------LeetCode Problem 20------------------------\n")

	for _, q := range qs {
		a, p := q.ans20, q.para20
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isValid(p.s), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}
