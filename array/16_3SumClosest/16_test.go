package leetcode

import (
	"fmt"
	"testing"
)

type question16 struct {
	para16
	ans16
}

type para16 struct {
	nums []int
	target int
}

type ans16 struct {
	ans int
}

func Test_Problem16(t *testing.T) {
	qs := []question16{{
		para16: para16{[]int{-1, 2, 1, -4}, 1},
		ans16:  ans16{2},
	}, {
		para16: para16{[]int{0, 2, 1, -3}, 1},
		ans16:  ans16{0},
	}, {
		para16: para16{[]int{6,-18,-20,-7,-15,9,18,10,1,-20,-17,-19,-3,-5,-19,10,6,-11,1,-17,-15,6,17,-18,-3,16,19,-20,-3,-17,-15,-3,12,1,-9,4,1,12,-2,14,4,-4,19,-20,6,0,-19,18,14,1,-15,-5,14,12,-4,0,-10,6,6,-6,20,-8,-6,5,0,3,10,7,-2,17,20,12,19,-13,-1,10,-1,14,0,7,-3,10,14,14,11,0,-4,-15,-8,3,2,-5,9,10,16,-4,-3,-9,-8,-14,10,6,2,-12,-7,-16,-6,10}, -52},
		ans16:  ans16{-52},
	}}

	fmt.Printf("------------------------Leetcode Problem 16------------------------\n")

	for _, q := range qs {
		a, p := q.ans16, q.para16
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, threeSumClosest(p.nums, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}