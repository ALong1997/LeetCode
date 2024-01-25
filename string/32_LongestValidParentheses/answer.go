package LeetCode

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/

/*
解法: 栈 括号匹配


结果: 执行用时 :0 ms 内存消耗 :3.5 MB
*/

func longestValidParentheses(s string) int {
	var stack []int
	stack = append(stack, -1) // 添加一个-1索引到栈里面，用于全部合法时计算最大长度
	ans := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			// 右括号弹出索引，再看上一个左括号的索引位置，用于计算最大长度
			// 如果栈为空,右括号索引进栈（这个实际用于占坑，不会用其索引）
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			// 当前元素的索引与栈顶元素作差，获取最近的括号匹配数
			if ans < i-stack[len(stack)-1] {
				ans = i - stack[len(stack)-1]
			}
		}
	}
	return ans
}
