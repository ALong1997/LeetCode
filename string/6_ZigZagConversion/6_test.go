package LeetCode

import (
	"fmt"
	"testing"
)

type question6 struct {
	para6
	ans6
}

type para6 struct {
	s       string
	numRows int
}

type ans6 struct {
	s string
}

func Test_Problem6(t *testing.T) {
	qs := []question6{{
		para6: para6{"LeetCodeISHIRING", 3},
		ans6:  ans6{"LCIRETOESIIGEDHN"},
	}, {
		para6: para6{"LeetCodeISHIRING", 4},
		ans6:  ans6{"LDREOEIIECIHNTSG"},
	}, {
		para6: para6{"PAYPALISHIRING", 4},
		ans6:  ans6{"PINALSIGYAHRPI"},
	}, {
		para6: para6{"AB", 1},
		ans6:  ans6{"AB"},
	}}

	fmt.Printf("------------------------LeetCode Problem 6------------------------\n")

	for _, q := range qs {
		a, p := q.ans6, q.para6
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, convert(p.s, p.numRows), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceConvert(p.s, p.numRows), a)
	}
	fmt.Printf("\n")
}
