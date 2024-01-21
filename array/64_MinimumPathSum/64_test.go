package leetcode

import (
	"fmt"
	"testing"
)

type question64 struct {
	para64
	ans64
}

type para64 struct {
	grid [][]int
}

type ans64 struct {
	ans int
}

func Test_Problem64(t *testing.T) {
	qs := []question64{{
		para64: para64{[][]int{{1,3,1},{1,5,1},{4,2,1}}},
		ans64:  ans64{7},
	}}

	fmt.Printf("------------------------Leetcode Problem 64------------------------\n")

	for _, q := range qs {
		a, p := q.ans64, q.para64
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, minPathSum(p.grid), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePathsWithObstacles(p.obstacleGrid), a)
	}
	fmt.Printf("\n")
}
