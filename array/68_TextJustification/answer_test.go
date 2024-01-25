package LeetCode

import (
	"fmt"
	"testing"
)

type question68 struct {
	para68
	ans68
}

type para68 struct {
	words    []string
	maxWidth int
}

type ans68 struct {
	ans []string
}

func Test_Problem68(t *testing.T) {
	qs := []question68{{
		para68: para68{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
		ans68: ans68{[]string{
			"This    is    an",
			"example  of text",
			"justification.  "}},
	}, {
		para68: para68{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
		ans68: ans68{[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        "}},
	}, {
		para68: para68{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
			"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
		ans68: ans68{[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  "}},
	}}

	fmt.Printf("------------------------LeetCode Problem 68------------------------\n")

	for _, q := range qs {
		a, p := q.ans68, q.para68
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, fullJustify(p.words, p.maxWidth), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceIsNumber(p.s), a)
	}
	fmt.Printf("\n")
}
