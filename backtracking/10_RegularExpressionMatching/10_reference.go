package LeetCode

/*
解法: 动态规划
因为题目拥有 最优子结构 ，一个自然的想法是将中间结果保存起来。
我们通过用 dp(i,j) 表示 text[i:] 和 pattern[j:] 是否能匹配。
我们可以用更短的字符串匹配问题来表示原本的问题。

我们用 myAnswer 中同样的回溯方法。
除此之外，因为函数 match(text[i:], pattern[j:]) 只会被调用一次，
我们用 dp(i, j) 来应对剩余相同参数的函数调用，这帮助我们节省了字符串建立操作所需要的时间，也让我们可以将中间结果进行保存。


结果: 执行用时 :4 ms 内存消耗 :2.9 MB
*/

// 自顶向下方法
var memo [][]int

// 1： true
// 0：未计算过
// -1：false
func reference1IsMatch(s string, p string) bool {
	memo = make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		memo[i] = make([]int, len(p)+1)
	}
	return dp(0, 0, s, p)
}

func dp(i, j int, s, p string) bool {
	if memo[i][j] != 0 {
		return memo[i][j] == 1
	}
	flag := false
	if j == len(p) {
		flag = i == len(s)
	} else {
		firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			flag = dp(i, j+2, s, p) || firstMatch && dp(i+1, j, s, p)
		} else {
			flag = firstMatch && dp(i+1, j+1, s, p)
		}
	}
	memo[i][j] = -1
	if flag {
		memo[i][j] = 1
	}
	return flag
}

// 自底向上方法
func reference2IsMatch(s string, p string) bool {
	memo := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		memo[i] = make([]bool, len(p)+1)
	}
	memo[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				memo[i][j] = memo[i][j+2] || firstMatch && memo[i+1][j]
			} else {
				memo[i][j] = firstMatch && memo[i+1][j+1]
			}
		}
	}
	return memo[0][0]
}
