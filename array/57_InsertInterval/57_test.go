package leetcode

import (
	"fmt"
	"testing"
)

type question57 struct {
	para57
	ans57
}

type para57 struct {
	intervals [][]int
	newInterval []int
}

type ans57 struct {
	ans [][]int
}

func Test_Problem57(t *testing.T) {
	qs := []question57{{
		para57: para57{[][]int{{1,3},{6,9}},[]int{2,5}},
		ans57:  ans57{[][]int{{1,5},{6,9}}},
	}, {
		para57: para57{[][]int{{1,2},{3,5},{6,7},{8,10},{12,16}},[]int{4,8}},
		ans57:  ans57{[][]int{{1,2},{3,10},{12,16}}},
	}, {
		para57: para57{[][]int{{3,5},{12,15}},[]int{6,6}},
		ans57:  ans57{[][]int{{3,5},{6,6},{12,15}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 57------------------------\n")

	for _, q := range qs {
		a, p := q.ans57, q.para57
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, insert(p.intervals, p.newInterval), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceCanJump(p.nums), a)
	}
	fmt.Printf("\n")
}
