package leetcode

import (
	"fmt"
	"testing"
)

type question47 struct {
	para47
	ans47
}

type para47 struct {
	nums []int
}

type ans47 struct {
	ans [][]int
}

func Test_Problem47(t *testing.T) {
	qs := []question47{{
		para47: para47{[]int{1,1,2}},
		ans47:  ans47{[][]int{{1,1,2},{1,2,1},{2,1,1}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 47------------------------\n")

	for _, q := range qs {
		a, p := q.ans47, q.para47
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, permuteUniqu(p.nums), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}