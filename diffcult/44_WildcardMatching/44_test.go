package leetcode

import (
	"fmt"
	"testing"
)

type question44 struct {
	para44
	ans44
}

type para44 struct {
	s string
	p string
}

type ans44 struct {
	ans bool
}

func Test_Problem42(t *testing.T) {
	qs := []question44{{
		para44: para44{"aa", "a"},
		ans44:  ans44{false},
	}, {
		para44: para44{"aa", "*"},
		ans44:  ans44{true},
	}, {
		para44: para44{"cb", "?a"},
		ans44:  ans44{false},
	}, {
		para44: para44{"adceb", "*a*b"},
		ans44:  ans44{true},
	}, {
		para44: para44{"acdcb", "a*c?b"},
		ans44:  ans44{false},
	}}

	fmt.Printf("------------------------Leetcode Problem 44------------------------\n")

	for _, q := range qs {
		a, p := q.ans44, q.para44
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isMatch(p.s, p.p), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2Trap(p.height), a)
	}
	fmt.Printf("\n")
}