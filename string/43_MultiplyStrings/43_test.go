package leetcode

import (
	"fmt"
	"testing"
)

type question43 struct {
	para43
	ans43
}

type para43 struct {
	nums1 string
	nums2 string
}

type ans43 struct {
	ans string
}

func Test_Problem43(t *testing.T) {
	qs := []question43{{
		para43: para43{"2", "3"},
		ans43:  ans43{"6"},
	}, {
		para43: para43{"123", "456"},
		ans43:  ans43{"56088"},
	}}

	fmt.Printf("------------------------Leetcode Problem 43------------------------\n")

	for _, q := range qs {
		a, p := q.ans43, q.para43
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, multiply(p.nums1,p.nums2), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2Trap(p.height), a)
	}
	fmt.Printf("\n")
}