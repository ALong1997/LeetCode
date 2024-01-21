package LeetCode

import (
	"fmt"
	"testing"
)

type question88 struct {
	para88
	ans88
}

type para88 struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

type ans88 struct {
	ans []int
}

func Test_Problem88(t *testing.T) {
	qs := []question88{{
		para88: para88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		ans88:  ans88{[]int{1, 2, 2, 3, 5, 6}},
	}, {
		para88: para88{[]int{0}, 0, []int{1}, 1},
		ans88:  ans88{[]int{1}},
	}}

	fmt.Printf("------------------------LeetCode Problem 88------------------------\n")

	for _, q := range qs {
		a, p := q.ans88, q.para88
		fmt.Printf("【input】:%v", p)
		merge(p.nums1, p.m, p.nums2, p.n)
		fmt.Printf("       【output】:%v       【answer】:%v\n", p.nums1, a)
	}
	fmt.Printf("\n")
}
