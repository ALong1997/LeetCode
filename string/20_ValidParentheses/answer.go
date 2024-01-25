package LeetCode

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
1.左括号必须用相同类型的右括号闭合。
2.左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/

/*
解法: 模拟栈
注意题目中有三种括号组合，先设置一个对应关系 Map
左括号类型入栈，右括号类型匹配栈顶元素，如果不是对应关系，则肯定 false，匹配上之后移除栈顶元素
最后看栈是否为空，空则是有效括号，非空则是无效的
同时注意题目中空字符串是返回 true

结果: 执行用时 :0 ms 内存消耗 :2 MB
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []byte

	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if m[s[i]] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
