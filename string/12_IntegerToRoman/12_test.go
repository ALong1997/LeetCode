package leetcode

import (
	"fmt"
	"testing"
)

type question12 struct {
	para12
	ans12
}

type para12 struct {
	num	int
}

type ans12 struct {
	ans string
}

func Test_Problem12(t *testing.T) {
	qs := []question12{{
		para12: para12{3},
		ans12:  ans12{"III"},
	}, {
		para12: para12{4},
		ans12:  ans12{"IV"},
	}, {
		para12: para12{9},
		ans12:  ans12{"IX"},
	}, {
		para12: para12{58},
		ans12:  ans12{"LVIII"},
	}, {
		para12: para12{1994},
		ans12:  ans12{"MCMXCIV"},
	}}

	fmt.Printf("------------------------Leetcode Problem 12------------------------\n")

	for _, q := range qs {
		a, p := q.ans12, q.para12
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, intToRoman(p.num), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}