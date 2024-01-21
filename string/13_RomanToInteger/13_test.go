package LeetCode

import (
	"fmt"
	"testing"
)

type question13 struct {
	para13
	ans13
}

type para13 struct {
	s string
}

type ans13 struct {
	ans int
}

func Test_Problem13(t *testing.T) {
	qs := []question13{{
		para13: para13{"III"},
		ans13:  ans13{3},
	}, {
		para13: para13{"IV"},
		ans13:  ans13{4},
	}, {
		para13: para13{"IX"},
		ans13:  ans13{9},
	}, {
		para13: para13{"LVIII"},
		ans13:  ans13{58},
	}, {
		para13: para13{"MCMXCIV"},
		ans13:  ans13{1994},
	}}

	fmt.Printf("------------------------LeetCode Problem 13------------------------\n")

	for _, q := range qs {
		a, p := q.ans13, q.para13
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, romanToInt(p.s), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}
