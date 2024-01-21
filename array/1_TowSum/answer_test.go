package LeetCode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	ans []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{{
		para1: para1{[]int{2, 7, 11, 15}, 9},
		ans1:  ans1{[]int{0, 1}},
	}, {
		para1: para1{[]int{3, 2, 4}, 6},
		ans1:  ans1{[]int{1, 2}},
	}, {
		para1: para1{[]int{3, 3}, 6},
		ans1:  ans1{[]int{0, 1}},
	}, {
		para1: para1{[]int{-1, -2, -3, -4, -5}, -8},
		ans1:  ans1{[]int{2, 4}},
	}, {
		para1: para1{[]int{2, 5, 5, 11}, 10},
		ans1:  ans1{[]int{1, 2}},
	}}

	fmt.Printf("------------------------LeetCode Problem 1------------------------\n")

	for _, q := range qs {
		a, p := q.ans1, q.para1
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, twoSum(p.nums, p.target), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, twoSum(p.nums, p.target), a)
	}
	fmt.Printf("\n")
}
