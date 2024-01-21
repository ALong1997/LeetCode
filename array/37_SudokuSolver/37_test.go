package leetcode

import (
	"fmt"
	"testing"
)

type question37 struct {
	para37
	ans37
}

type para37 struct {
	board [][]byte
}

type ans37 struct {

}

func Test_Problem37(t *testing.T) {
	qs := []question37{{
		para37: para37{[][]byte{
			{'5','3','.','.','7','.','.','.','.',},
			{'6','.','.','1','9','5','.','.','.',},
			{'.','9','8','.','.','.','.','6','.',},
			{'8','.','.','.','6','.','.','.','3',},
			{'4','.','.','8','.','3','.','.','1',},
			{'7','.','.','.','2','.','.','.','6',},
			{'.','6','.','.','.','.','2','8','.',},
			{'.','.','.','4','1','9','.','.','5',},
			{'.','.','.','.','8','.','.','7','9',},
		}},
	}}

	fmt.Printf("------------------------Leetcode Problem 37------------------------\n")

	for _, q := range qs {
		a, p := q.ans37, q.para37
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, nil, a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsValidSudoku(p.board), a)
	}
	fmt.Printf("\n")
}