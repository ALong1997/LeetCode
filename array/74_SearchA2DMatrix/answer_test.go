package LeetCode

import (
	"fmt"
	"testing"
)

type question74 struct {
	para74
	ans74
}

type para74 struct {
	matrix [][]int
	target int
}

type ans74 struct {
	ans bool
}

func Test_Problem74(t *testing.T) {
	qs := []question74{{
		para74: para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3},
		ans74:  ans74{true},
	}, {
		para74: para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13},
		ans74:  ans74{false},
	}, {
		para74: para74{[][]int{{1, 1}}, 2},
		ans74:  ans74{false},
	}}

	fmt.Printf("------------------------LeetCode Problem 74------------------------\n")

	for _, q := range qs {
		a, p := q.ans74, q.para74
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, searchMatrix(p.matrix, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePathsWithObstacles(p.obstacleGrid), a)
	}
	fmt.Printf("\n")
}
