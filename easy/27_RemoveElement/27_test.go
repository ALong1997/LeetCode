package leetcode

import (
	"fmt"
	"testing"
)

type question27 struct {
	para27
	ans27
}

type para27 struct {
	nums []int
	val int
}

type ans27 struct {
	ans int
}

func Test_Problem27(t *testing.T) {
	qs := []question27{{
		para27: para27{[]int{3, 2, 2, 3}, 3},
		ans27:  ans27{2},
	}, {
		para27: para27{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
		ans27:  ans27{5},
	}}

	fmt.Printf("------------------------Leetcode Problem 13------------------------\n")

	for _, q := range qs {
		a, p := q.ans27, q.para27
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, removeElement(p.nums, p.val), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}