package leetcode

/*
解法: 动态规划
参考 https://www.jianshu.com/p/2eeb6599ec30


结果: 执行用时 :20 ms 内存消耗 :6.1 MB
*/


func referenceIsMatch(s string, p string) bool {
	var col, row = len(s), len(p)
	//建立一个boolean数组，数组长度为 row + 1 * col + 1，多的第一列第一行代表空串的情况
	var dp = make([][]bool, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]bool, col+1)
	}

	// 初始化第一行，第一行剩余元素全部变为false，因为pattern如果是空串，那么只要s不是空串，都不能匹配
	dp[0][0] = true
	//for i := 1; i < col+1; i++ {
	//	dp[0][i] = false
	//}
	//初始化第0列，此时s是空串，所以只能是*(***)这种形式的
	//注意，i-1代表的真实的p的索引
	for i := 1; i < row+1; i++ {
		if p[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = false
		}
	}

	for i := 1; i < col+1; i++ {
		for j := 1; j < row+1; j++ {
			if p[j-1] == '*' {
				//要么是上面是true 要么是左边是true，如果上面是true，此时*当空串使，如果左边是true，那么此时*当一个字符使
				if dp[j-1][i] || dp[j][i-1] {
					dp[j][i] = true
				} else {
					dp[j][i] = false
				}
			} else {
				//如果不为*的话，首先左上角的位置要是true，如果不是的话，s和p都增加一个字符也不会匹配，同时s和p增加的字符要相同或者p是？
				if dp[j-1][i-1] && (p[j-1] == s[i-1] || p[j-1] == '?') {
					dp[j][i] = true
				} else {
					dp[j][i] = false
				}
			}
		}
	}
	return dp[row][col]
}