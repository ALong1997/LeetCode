package leetcode

import (
	"fmt"
	"testing"
)

type question28 struct {
	para28
	ans28
}

type para28 struct {
	aystack string
	needle string
}

type ans28 struct {
	ans int
}

func Test_Problem28(t *testing.T) {
	qs := []question28{{
		para28: para28{"hello", "ll"},
		ans28:  ans28{2},
	}, {
		para28: para28{"aaaab", "bba"},
		ans28:  ans28{-1},
	}, {
		para28: para28{"", "a"},
		ans28:  ans28{-1},
	}}

	fmt.Printf("------------------------Leetcode Problem 28------------------------\n")

	for _, q := range qs {
		a, p := q.ans28, q.para28
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, strStr(p.aystack, p.needle), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceStrStr(p.aystack, p.needle), a)
	}
	fmt.Printf("\n")
}