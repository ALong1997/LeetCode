package LeetCode

import (
	"fmt"
	"testing"
)

type question67 struct {
	para67
	ans67
}

type para67 struct {
	a string
	b string
}

type ans67 struct {
	ans string
}

func Test_Problem67(t *testing.T) {
	qs := []question67{{
		para67: para67{"11", "1"},
		ans67:  ans67{"100"},
	}, {
		para67: para67{"1010", "1011"},
		ans67:  ans67{"10101"},
	}}

	fmt.Printf("------------------------LeetCode Problem 67------------------------\n")

	for _, q := range qs {
		a, p := q.ans67, q.para67
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, addBinary(p.a, p.b), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
