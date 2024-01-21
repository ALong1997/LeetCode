package LeetCode

import (
	"fmt"
	"testing"
)

type question36 struct {
	para36
	ans36
}

type para36 struct {
	board [][]byte
}

type ans36 struct {
	ans bool
}

func Test_Problem36(t *testing.T) {
	qs := []question36{{
		para36: para36{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}},
		ans36: ans36{true},
	}}

	fmt.Printf("------------------------LeetCode Problem 36------------------------\n")

	for _, q := range qs {
		a, p := q.ans36, q.para36
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isValidSudoku(p.board), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsValidSudoku(p.board), a)
	}
	fmt.Printf("\n")
}
