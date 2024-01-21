package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

type para17 struct {
	digits string
}

type ans17 struct {
	ans []string
}

func Test_Problem17(t *testing.T) {
	qs := []question17{{
		para17: para17{"23"},
		ans17:  ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}}

	fmt.Printf("------------------------Leetcode Problem 17------------------------\n")

	for _, q := range qs {
		a, p := q.ans17, q.para17
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, letterCombinations(p.digits), a)
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceThreeSum(p.nums), a)
	}
	fmt.Printf("\n")
}