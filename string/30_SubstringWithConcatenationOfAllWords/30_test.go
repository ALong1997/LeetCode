package leetcode

import (
	"fmt"
	"testing"
)

type question30 struct {
	para30
	ans30
}

type para30 struct {
	s string
	words []string
}

type ans30 struct {
	ans []int
}

func Test_Problem30(t *testing.T) {
	qs := []question30{{
		para30: para30{"barfoothefoobarman", []string{"foo","bar"}},
		ans30:  ans30{[]int{0, 9}},
	}, {
		para30: para30{"wordgoodgoodgoodbestword", []string{"word","good","best","word"}},
		ans30:  ans30{[]int{}},
	}}

	fmt.Printf("------------------------Leetcode Problem 30------------------------\n")

	for _, q := range qs {
		a, p := q.ans30, q.para30
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, findSubstring(p.s, p.words), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceStrStr(p.aystack, p.needle), a)
	}
	fmt.Printf("\n")
}