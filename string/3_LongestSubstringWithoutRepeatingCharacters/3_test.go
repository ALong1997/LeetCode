package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	para3
	ans3
}

type para3 struct {
	s		string
}

type ans3 struct {
	length int
}

func Test_Problem3(t *testing.T)  {
	qs := []question3{{
			para3: para3{"abcabcbb"},
			ans3:  ans3{3},
		}, {
			para3: para3{"bbbbb"},
			ans3:  ans3{1},
		}, {
			para3: para3{"pwwkew"},
			ans3:  ans3{3},
		}}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, q := range qs {
		a, p := q.ans3, q.para3
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, lengthOfLongestSubstring(p.s), a)
	}
	fmt.Printf("\n")
}