package leetcode

import (
	"fmt"
	"testing"
)

type question85 struct {
	para85
	ans85
}

type para85 struct {
	matrix [][]byte
}

type ans85 struct {
	ans int
}

func Test_Problem85(t *testing.T) {
	qs := []question85{{
		para85: para85{[][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}},
		ans85:  ans85{6},
	}}

	fmt.Printf("------------------------Leetcode Problem 85------------------------\n")

	for _, q := range qs {
		a, p := q.ans85, q.para85
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, maximalRectangle(p.matrix), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, largestRectangleAreaHeight(p.heights), a)
	}
	fmt.Printf("\n")
}
