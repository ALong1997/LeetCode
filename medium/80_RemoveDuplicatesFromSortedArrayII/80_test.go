package leetcode

import (
	"fmt"
	"testing"
)

type question80 struct {
	para80
	ans80
}

type para80 struct {
	nums []int
}

type ans80 struct {
	ans int
}

func Test_Problem80(t *testing.T) {
	qs := []question80{{
		para80: para80{[]int{1, 1, 1, 2, 2, 3}},
		ans80:  ans80{5},
	}, {
		para80: para80{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
		ans80:  ans80{7},
	}, {
		para80: para80{[]int{1, 1, 1}},
		ans80:  ans80{2},
	}, {
		para80: para80{[]int{0, 0, 0, 0, 0, 2}},
		ans80:  ans80{3},
	}}

	fmt.Printf("------------------------Leetcode Problem 80------------------------\n")

	for _, q := range qs {
		a, p := q.ans80, q.para80
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, removeDuplicates(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceSubsets(p.nums), a)
	}
	fmt.Printf("\n")
}
