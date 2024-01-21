package leetcode

import (
	"fmt"
	"testing"
)

type question31 struct {
	para31
	ans31
}

type para31 struct {
	nums []int
}

type ans31 struct {

}

func Test_Problem31(t *testing.T) {
	qs := []question31{{
		para31: para31{[]int{1, 2, 3}},
		ans31:  ans31{},
	}, {
		para31: para31{[]int{3, 2, 1}},
		ans31:  ans31{},
	}}

	fmt.Printf("------------------------Leetcode Problem 31------------------------\n")

	for _, q := range qs {
		a, p := q.ans31, q.para31
		fmt.Printf("【input】:%v				", p)
		nextPermutation(p.nums)
		fmt.Printf("【output】:%v				【answer】:%v\n", p, a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceStrStr(p.aystack, p.needle), a)
	}
	fmt.Printf("\n")
}