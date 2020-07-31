package leetcode

import (
	"fmt"
	"testing"
)

type question72 struct {
	para72
	ans72
}

type para72 struct {
	words1 string
	words2 string
}

type ans72 struct {
	ans int
}

func Test_Problem72(t *testing.T) {
	qs := []question72{{
		para72: para72{"horse", "ros"},
		ans72:  ans72{3},
	}, {
		para72: para72{"intention", "execution"},
		ans72:  ans72{5},
	}}

	fmt.Printf("------------------------Leetcode Problem 72------------------------\n")

	for _, q := range qs {
		a, p := q.ans72, q.para72
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, minDistance(p.words1, p.words2), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
