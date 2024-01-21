package leetcode

import (
	"fmt"
	"testing"
)

type question10 struct {
	para10
	ans10
}

type para10 struct {
	s	string
	p	string
}

type ans10 struct {
	ans bool
}

func Test_Problem9(t *testing.T) {
	qs := []question10{{
		para10: para10{"aa", "a"},
		ans10:  ans10{false},
	}, {
		para10: para10{"aa", "a*"},
		ans10:  ans10{true},
	}, {
		para10: para10{"aab", "c*a*b"},
		ans10:  ans10{true},
	}, {
		para10: para10{"mississippi", "mis*is*p*."},
		ans10:  ans10{false},
	}, {
		para10: para10{"ab", ".*"},
		ans10:  ans10{true},
	}, {
		para10: para10{"aaa", "a*a"},
		ans10:  ans10{true},
	}, {
		para10: para10{"aaa", "ab*a*c*a"},
		ans10:  ans10{true},
	}, {
		para10: para10{"ssip", "s*p*"},
		ans10:  ans10{false},
	}}

	fmt.Printf("------------------------Leetcode Problem 5------------------------\n")

	for _, q := range qs {
		a, p := q.ans10, q.para10
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isMatch(p.s, p.p), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference1IsMatch(p.s, p.p), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference3IsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}