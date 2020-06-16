package leetcode

import (
	"fmt"
	"testing"
)

type question54 struct {
	para54
	ans54
}

type para54 struct {
	matrix [][]int
}

type ans54 struct {
	ans []int
}

func Test_Problem54(t *testing.T) {
	qs := []question54{{
		para54: para54{[][]int{{1,2,3},{4,5,6},{7,8,9}}},
		ans54:  ans54{[]int{1,2,3,6,9,8,7,4,5}},
	}, {
		para54: para54{[][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}},
		ans54:  ans54{[]int{1,2,3,4,8,12,11,10,9,5,6,7}},
	}}

	fmt.Printf("------------------------Leetcode Problem 54------------------------\n")

	for _, q := range qs {
		a, p := q.ans54, q.para54
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, spiralOrder(p.matrix), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceSpiralOrder(p.matrix), a)
	}
	fmt.Printf("\n")
}
