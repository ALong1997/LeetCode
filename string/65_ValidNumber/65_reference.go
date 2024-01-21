package leetcode

import (
	"regexp"
	"strings"
)

/*
解法: 正则表达式


结果: 执行用时 :20 ms 内存消耗 :6.7 MB
*/

func referenceIsNumber(s string) bool {
	s = strings.TrimSpace(s)
	ans,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)

	return ans
}
