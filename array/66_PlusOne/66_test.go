package leetcode

import (
	"fmt"
	"testing"
)

type question66 struct {
	para66
	ans66
}

type para66 struct {
	digits []int
}

type ans66 struct {
	ans []int
}

func Test_Problem66(t *testing.T) {
	qs := []question66{{
		para66: para66{[]int{1,2,3}},
		ans66:  ans66{[]int{1,2,4}},
	}, {
		para66: para66{[]int{4,3,2,1}},
		ans66:  ans66{[]int{4,3,2,2}},
	}, {
		para66: para66{[]int{9}},
		ans66:  ans66{[]int{1,0}},
	}}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		a, p := q.ans66, q.para66
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, plusOne(p.digits), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
