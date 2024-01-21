package leetcode

import (
	"fmt"
	"testing"
)

type question76 struct {
	para76
	ans76
}

type para76 struct {
	s string
	t string
}

type ans76 struct {
	ans string
}

func Test_Problem76(t *testing.T) {
	qs := []question76{{
		para76: para76{"ADOBECODEBANC", "ABC"},
		ans76:  ans76{"BANC"},
	}}

	fmt.Printf("------------------------Leetcode Problem 76------------------------\n")

	for _, q := range qs {
		a, p := q.ans76, q.para76
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, minWindow(p.s, p.t), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
