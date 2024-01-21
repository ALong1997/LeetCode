package LeetCode

import (
	"fmt"
	"testing"
)

type question40 struct {
	para40
	ans40
}

type para40 struct {
	candidates []int
	target     int
}

type ans40 struct {
	ans [][]int
}

func Test_Problem40(t *testing.T) {
	qs := []question40{{
		para40: para40{[]int{2, 3, 5}, 8},
		ans40:  ans40{[][]int{{3, 5}}},
	}, {
		para40: para40{[]int{2, 3, 5}, 7},
		ans40:  ans40{[][]int{{2, 5}}},
	}, {
		para40: para40{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
		ans40:  ans40{[][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}},
	}, {
		para40: para40{[]int{2, 5, 2, 1, 2}, 5},
		ans40:  ans40{[][]int{{1, 2, 2}, {5}}},
	}}

	fmt.Printf("------------------------LeetCode Problem 40------------------------\n")

	for _, q := range qs {
		a, p := q.ans40, q.para40
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, combinationSum2(p.candidates, p.target), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsValidSudoku(p.board), a)
	}
	fmt.Printf("\n")
}
