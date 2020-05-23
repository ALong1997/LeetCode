package leetcode

/*
解法: 动态规划
我们定义一个 dp 数组，其中第 i 个元素表示以下标为 i 的字符结尾的最长有效子字符串的长度。
我们将 dp 数组全部初始化为 0 。现在，很明显有效的子字符串一定以 ‘)’ 结尾。这进一步可以得出结论：以 ‘(’ 结尾的子字符串对应的 dp 数组位置上的值必定为 0 。
所以说我们只需要更新 ‘)’ 在 dp 数组中对应位置的值。

那么分为几种情况:
        1) 当前字符为')'时,一定构不成有效子串直接忽略
        2) 当前字符为'('时
            a.s[i-1] 为'('
                则 i==0时 dp[i] = 2
                   i>0时  dp[i] = dp[i-1] + 2
            b. s[i-1] 为'('
                则  i - dp[i - 1] - 1 == 0 时
                      dp[i] = dp[i-1] + 2
                    i - dp[i - 1] - 1 != 0 时
                      dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]


结果: 执行用时 :0 ms 内存消耗 :2.8 MB
*/


func referenceLongestValidParentheses(s string) int {
	dp, ans := make([]int,len(s)), 0
	// dp 数组表示 当前位置i处 之前在s串中构造的有效括号长度
	for i,_ := range s{
		if i > 0 &&  s[i] == ')'{
			if s[i - 1] == '('{
				if i == 1{
					dp[i] = 2
				}else{
					dp [i] = dp[i - 2] + 2
				}
			}else if s[i - 1] == ')' &&  i - dp[i - 1] - 1 >= 0   &&  s[i-dp[i-1]-1] == '('  {
				if  i - dp[i - 1] - 1 == 0{
					dp[i] = dp[i-1] + 2
				}else{
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > ans{
				ans = dp[i]
			}
		}
	}
	return ans
}