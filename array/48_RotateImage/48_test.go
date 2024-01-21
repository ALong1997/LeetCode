package LeetCode

import (
	"fmt"
	"testing"
)

type question48 struct {
	para48
	ans48
}

type para48 struct {
	matrix [][]int
}

type ans48 struct {
}

func Test_Problem48(t *testing.T) {
	qs := []question48{{
		para48: para48{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		ans48:  ans48{},
	}}

	fmt.Printf("------------------------LeetCode Problem 48------------------------\n")

	for _, q := range qs {
		a, p := q.ans48, q.para48
		fmt.Printf("【input】				:%v", p)
		rotate(p.matrix)
		fmt.Printf("【output】:%v				【answer】:%v\n", p, a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsMatch(p.s, p.p), a)
	}
	fmt.Printf("\n")
}
