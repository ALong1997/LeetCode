package leetcode

import (
	"fmt"
	"testing"
)

type question41 struct {
	para41
	ans41
}

type para41 struct {
	nums []int
}

type ans41 struct {
	ans int
}

func Test_Problem41(t *testing.T) {
	qs := []question41{{
		para41: para41{[]int{1,2,0}},
		ans41:  ans41{3},
	}, {
		para41: para41{[]int{3,4,-1,1}},
		ans41:  ans41{2},
	}, {
		para41: para41{[]int{7,8,9,11,12}},
		ans41:  ans41{1},
	}}

	fmt.Printf("------------------------Leetcode Problem 41------------------------\n")

	for _, q := range qs {
		a, p := q.ans41, q.para41
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, firstMissingPositive(p.nums), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceFirstMissingPositive(p.nums), a)
	}
	fmt.Printf("\n")
}