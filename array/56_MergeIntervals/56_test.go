package leetcode

import (
	"fmt"
	"testing"
)

type question56 struct {
	para56
	ans56
}

type para56 struct {
	intervals [][]int
}

type ans56 struct {
	ans [][]int
}

func Test_Problem56(t *testing.T) {
	qs := []question56{{
		para56: para56{[][]int{{1,3},{2,6},{8,10},{15,18}}},
		ans56:  ans56{[][]int{{1,6},{8,10},{15,18}}},
	}, {
		para56: para56{[][]int{{1,4},{4,5}}},
		ans56:  ans56{[][]int{{1,5}}},
	}, {
		para56: para56{[][]int{{1,4},{2,3}}},
		ans56:  ans56{[][]int{{1,4}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 56------------------------\n")

	for _, q := range qs {
		a, p := q.ans56, q.para56
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, merge(p.intervals), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
