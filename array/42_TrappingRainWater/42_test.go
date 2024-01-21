package LeetCode

import (
	"fmt"
	"testing"
)

type question42 struct {
	para42
	ans42
}

type para42 struct {
	height []int
}

type ans42 struct {
	ans int
}

func Test_Problem42(t *testing.T) {
	qs := []question42{{
		para42: para42{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		ans42:  ans42{6},
	}}

	fmt.Printf("------------------------LeetCode Problem 42------------------------\n")

	for _, q := range qs {
		a, p := q.ans42, q.para42
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, trap(p.height), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, reference2Trap(p.height), a)
	}
	fmt.Printf("\n")
}
