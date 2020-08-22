package leetcode

import (
	"fmt"
	"testing"
)

type question79 struct {
	para79
	ans79
}

type para79 struct {
	board [][]byte
	word  string
}

type ans79 struct {
	ans bool
}

func Test_Problem79(t *testing.T) {
	qs := []question79{{
		para79: para79{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'}},
			"ABCCED"},
		ans79: ans79{true},
	}, {
		para79: para79{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'}},
			"SEE"},
		ans79: ans79{true},
	}, {
		para79: para79{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'}},
			"ABCB"},
		ans79: ans79{false},
	}}

	fmt.Printf("------------------------Leetcode Problem 79------------------------\n")

	for _, q := range qs {
		a, p := q.ans79, q.para79
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, exist(p.board, p.word), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceSubsets(p.nums), a)
	}
	fmt.Printf("\n")
}
