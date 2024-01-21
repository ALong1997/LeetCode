package LeetCode

import (
	"fmt"
	"testing"
)

type question71 struct {
	para71
	ans71
}

type para71 struct {
	path string
}

type ans71 struct {
	ans string
}

func Test_Problem71(t *testing.T) {
	qs := []question71{{
		para71: para71{"/home/"},
		ans71:  ans71{"/home/"},
	}, {
		para71: para71{"/../"},
		ans71:  ans71{"/"},
	}, {
		para71: para71{"/home//foo/"},
		ans71:  ans71{"/home/foo"},
	}, {
		para71: para71{"/a/./b/../../c/"},
		ans71:  ans71{"/c"},
	}, {
		para71: para71{"/a/../../b/../c//.//"},
		ans71:  ans71{"/c"},
	}, {
		para71: para71{"/a//b////c/d//././/.."},
		ans71:  ans71{"/a/b/c"},
	}}

	fmt.Printf("------------------------LeetCode Problem 71------------------------\n")

	for _, q := range qs {
		a, p := q.ans71, q.para71
		fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, simplifyPath(p.path), a)
		//fmt.Printf("【input】:%v       【output】:%v       【answer】:%v\n", p, referenceUniquePathsWithObstacles(p.obstacleGrid), a)
	}
	fmt.Printf("\n")
}
