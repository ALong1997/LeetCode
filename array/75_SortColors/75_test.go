package leetcode

import (
	"fmt"
	"testing"
)

type question75 struct {
	para75
	ans75
}

type para75 struct {
	nums []int
}

type ans75 struct {
	ans []int
}

func Test_Problem75(t *testing.T) {
	qs := []question75{{
		para75: para75{[]int{2, 0, 2, 1, 1, 0}},
		ans75:  ans75{[]int{0, 0, 1, 1, 2, 2}},
	}}

	fmt.Printf("------------------------Leetcode Problem 75------------------------\n")

	for _, q := range qs {
		a, p := q.ans75, q.para75
		fmt.Printf("【input】:%v       ", p)
		sortColors(p.nums)
		fmt.Printf("【output】:%v       【answer】:%v\n", p, a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePathsWithObstacles(p.obstacleGrid), a)
	}
	fmt.Printf("\n")
}
