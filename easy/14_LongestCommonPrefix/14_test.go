package leetcode

import (
	"fmt"
	"testing"
)

type question14 struct {
	para14
	ans14
}

type para14 struct {
	strs []string
}

type ans14 struct {
	ans string
}

func Test_Problem14(t *testing.T) {
	qs := []question14{{
		para14: para14{[]string{"flower", "flow", "flight"}},
		ans14:  ans14{"fl"},
	},{
		para14: para14{[]string{"dog", "racecar", "car"}},
		ans14:  ans14{""},
	},{
		para14: para14{[]string{"a"}},
		ans14:  ans14{"a"},
	},{
		para14: para14{[]string{}},
		ans14:  ans14{""},
	}}

	fmt.Printf("------------------------Leetcode Problem 14------------------------\n")

	for _, q := range qs {
		a, p := q.ans14, q.para14
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, longestCommonPrefix(p.strs), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}