/*
验证给定的字符串是否可以解释为十进制数字。

说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
	数字 0-9
	指数 - "e"
	正/负号 - "+"/"-"
	小数点 - "."
	当然，在输入中，这些字符的上下文也很重要。
*/
package leetcode

import "strings"

/*
解法: 有限状态机


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

const (
	blank  = 0 // 空格
	digit1 = 1 // 数字(0-9) 无前缀
	sign1  = 2 // +/- 无e前缀
	point  = 4 // '.'
	digit2 = 5 // 数字(0-9) 有符号前缀
	e      = 6 // 'e'
	sign2  = 7 // +/- 有e前缀
	digit3 = 8 // 数字(0-9) 有e前缀
	err    = 9
)

func isNumber(s string) bool {
	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		[]int{blank, digit1, sign1, point, err},
		[]int{err, digit1, err, digit2, e},
		[]int{err, digit1, err, point, err},
		[]int{err, digit2, err, err, e},
		[]int{err, digit2, err, err, err},
		[]int{err, digit2, err, err, e},
		[]int{err, digit3, sign2, err, err},
		[]int{err, digit3, err, err, err},
		[]int{err, digit3, err, err, err},
	}

	state := 0 // blank start
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 1
		case '+', '-':
			newState = 2
		case '.':
			newState = 3
		case 'e':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == err {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}
