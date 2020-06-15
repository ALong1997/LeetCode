package leetcode

import (
	"fmt"
	"testing"
)

type question53 struct {
	para53
	ans53
}

type para53 struct {
	nums []int
}

type ans53 struct {
	ans int
}

func Test_Problem53(t *testing.T) {
	qs := []question53{{
		para53: para53{[]int{-2,1,-3,4,-1,2,1,-5,4}},
		ans53:  ans53{6},
	}}

	fmt.Printf("------------------------Leetcode Problem 53------------------------\n")

	for _, q := range qs {
		a, p := q.ans53, q.para53
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, maxSubArray(p.nums), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMaxSubArray(p.nums), a)
	}
	fmt.Printf("\n")
}
