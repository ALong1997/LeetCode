package leetcode

import (
	"fmt"
	"testing"
)

type question63 struct {
	para63
	ans63
}

type para63 struct {
	obstacleGrid [][]int
}

type ans63 struct {
	ans int
}

func Test_Problem63(t *testing.T) {
	qs := []question63{{
		para63: para63{[][]int{{0,0,0},{0,1,0},{0,0,0}}},
		ans63:  ans63{2},
	}, {
		para63: para63{[][]int{{0,0},{1,0}}},
		ans63:  ans63{1},
	}}

	fmt.Printf("------------------------Leetcode Problem 63------------------------\n")

	for _, q := range qs {
		a, p := q.ans63, q.para63
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, uniquePathsWithObstacles(p.obstacleGrid), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePaths(p.m, p.n), a)
	}
	fmt.Printf("\n")
}
