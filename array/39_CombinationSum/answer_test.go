package LeetCode

import (
	"fmt"
	"testing"
)

type question39 struct {
	para39
	ans39
}

type para39 struct {
	candidates []int
	target     int
}

type ans39 struct {
	ans [][]int
}

func Test_Problem39(t *testing.T) {
	qs := []question39{{
		para39: para39{[]int{2, 3, 6, 7}, 7},
		ans39:  ans39{[][]int{{7}, {2, 2, 3}}},
	}, {
		para39: para39{[]int{2, 3, 5}, 8},
		ans39:  ans39{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}, {
		para39: para39{[]int{2, 3, 5}, 7},
		ans39:  ans39{[][]int{{2, 5}, {2, 2, 3}}},
	}}

	fmt.Printf("------------------------LeetCode Problem 39------------------------\n")

	for _, q := range qs {
		a, p := q.ans39, q.para39
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, combinationSum(p.candidates, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsValidSudoku(p.board), a)
	}
	fmt.Printf("\n")
}
