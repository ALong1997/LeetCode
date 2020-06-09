package leetcode

import (
	"fmt"
	"testing"
)

type question49 struct {
	para49
	ans49
}

type para49 struct {
	str []string
}

type ans49 struct {
	ans [][]string
}

func Test_Problem49(t *testing.T) {
	qs := []question49{{
		para49: para49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		ans49:  ans49{[][]string{{"ate","eat","tea"}, {"nat","tan"}, {"bat"}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 49------------------------\n")

	for _, q := range qs {
		a, p := q.ans49, q.para49
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, groupAnagrams(p.str), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}