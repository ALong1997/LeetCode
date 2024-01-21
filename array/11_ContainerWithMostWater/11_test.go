package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct {
	para11
	ans11
}

type para11 struct {
	height []int
}

type ans11 struct {
	ans int
}

func Test_Problem11(t *testing.T) {
	qs := []question11{{
		para11: para11{[]int{1,8,6,2,5,4,8,3,7}},
		ans11:  ans11{49},
	}}

	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")

	for _, q := range qs {
		a, p := q.ans11, q.para11
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, maxArea(p.height), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}