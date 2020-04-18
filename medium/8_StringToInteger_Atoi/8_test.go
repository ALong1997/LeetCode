package leetcode

import (
	"fmt"
	"testing"
)

type question8 struct {
	para8
	ans8
}

type para8 struct {
	str	string
}

type ans8 struct {
	ans int
}

func Test_Problem8(t *testing.T) {
	qs := []question8{{
		para8: para8{"42"},
		ans8:  ans8{42},
	}, {
		para8: para8{"   -42"},
		ans8:  ans8{-42},
	}, {
		para8: para8{"4193 with words"},
		ans8:  ans8{4193},
	}, {
		para8: para8{"words and 987"},
		ans8:  ans8{0},
	}, {
		para8: para8{"-91283472332"},
		ans8:  ans8{-2147483648},
	}, {
		para8: para8{"000000000000000000"},
		ans8:  ans8{0},
	}, {
		para8: para8{"9223372036854775808"},
		ans8:  ans8{2147483647},
	}}

	fmt.Printf("------------------------Leetcode Problem 6------------------------\n")

	for _, q := range qs {
		a, p := q.ans8, q.para8
		// fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, myAtoi(p.str), a)
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceMyAtoi(p.str), a)
	}
	fmt.Printf("\n")
}