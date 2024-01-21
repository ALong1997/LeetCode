package leetcode

/*
解法:不需要额外的空间
在这种方法中，我们利用两个计数器 left 和 right 。
首先，我们从左到右遍历字符串，对于遇到的每个 ‘(’，我们增加 left 计算器，对于遇到的每个 ‘)’ ，我们增加 right 计数器。
每当 left 计数器与 right 计数器相等时，我们计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串。
如果 right 计数器比 eft 计数器大时，我们将 left 和 right 计数器同时变回 0 。
接下来，我们从右到左做一遍类似的工作。(原因：如果从左向右遍历，每当右括号多了的时候，栈就会为空，此时就可以统计之前的有效括号的长度。但是如果是左括号多了，那直到最后也不会栈空，此时也就无法统计中间是否存在有效括号了。因此，从左向右遍历时，一旦遇到右半部分「第一个多余的左括号」，那这个位置后面的串就都无法统计了。)

结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/


func reference2LongestValidParentheses(s string) int {
	var left, right, ans int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if ans < 2 * right {
				ans = 2 * right
			}
		} else if right >= left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if ans < 2 * left {
				ans = 2 * left
			}
		} else if left >= right {
			left, right = 0, 0
		}
	}
	return ans
}