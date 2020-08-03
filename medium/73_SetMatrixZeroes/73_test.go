package leetcode

import (
	"fmt"
	"testing"
)

type question73 struct {
	para73
	ans73
}

type para73 struct {
	matrix [][]int
}

type ans73 struct {
	matrix [][]int
}

func Test_Problem73(t *testing.T) {
	qs := []question73{{
		para73: para73{[][]int{{1,1,1},{1,0,1},{1,1,1}}},
		ans73:  ans73{[][]int{{1,0,1},{0,0,0},{1,0,1}}},
	}, {
		para73: para73{[][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}},
		ans73:  ans73{[][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}}},
	}}

	fmt.Printf("------------------------Leetcode Problem 73------------------------\n")

	for _, q := range qs {
		a, p := q.ans73, q.para73
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, p.matrix, a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePathsWithObstacles(p.obstacleGrid), a)
	}
	fmt.Printf("\n")
}
