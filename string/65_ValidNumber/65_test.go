package LeetCode

import (
	"fmt"
	"testing"
)

type question65 struct {
	para65
	ans65
}

type para65 struct {
	s string
}

type ans65 struct {
	ans bool
}

func Test_Problem65(t *testing.T) {
	qs := []question65{{
		para65: para65{"0"},
		ans65:  ans65{true},
	}, {
		para65: para65{"0.1"},
		ans65:  ans65{true},
	}, {
		para65: para65{"abc"},
		ans65:  ans65{false},
	}, {
		para65: para65{"1 a"},
		ans65:  ans65{false},
	}, {
		para65: para65{"2e10"},
		ans65:  ans65{true},
	}, {
		para65: para65{"-90e3"},
		ans65:  ans65{true},
	}, {
		para65: para65{"1e"},
		ans65:  ans65{false},
	}, {
		para65: para65{"e3"},
		ans65:  ans65{false},
	}, {
		para65: para65{"6e-1"},
		ans65:  ans65{true},
	}, {
		para65: para65{"99e2.5"},
		ans65:  ans65{true},
	}, {
		para65: para65{"53.5e93"},
		ans65:  ans65{true},
	}, {
		para65: para65{"--6"},
		ans65:  ans65{false},
	}, {
		para65: para65{"-+3"},
		ans65:  ans65{false},
	}, {
		para65: para65{"95a54e53"},
		ans65:  ans65{false},
	}, {
		para65: para65{"1 "},
		ans65:  ans65{true},
	}, {
		para65: para65{".1"},
		ans65:  ans65{true},
	}}

	fmt.Printf("------------------------LeetCode Problem 65------------------------\n")

	for _, q := range qs {
		a, p := q.ans65, q.para65
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, isNumber(p.s), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
