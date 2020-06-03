/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
*/
package leetcode

/*
解法: 参考T10


结果: 执行用时 :12 ms 内存消耗 :6.1 MB
*/


func isMatch(s string, p string) bool {
	var sLength, pLength = len(s), len(p)

	if p == s || p == "*" {
		return true
	}
	if p == "" || s == "" {
		return false
	}

	var dp = make([][]bool, pLength + 1)
	for i := 0; i < pLength + 1 ; i++ {
		dp[i] = make([]bool, sLength + 1)
	}
	dp[0][0] = true

	var sIndex, pIndex int
	for pIndex = 1; pIndex < pLength+1; pIndex++ {
		if p[pIndex-1] == '*' {
			// the current character in the pattern is '*'
			sIndex = 1
			// d[p_idx - 1][s_idx - 1] is a string-pattern match
			// on the previous step, i.e. one character before.
			// Find the first idx in string with the previous math.
			for !dp[pIndex-1][sIndex-1] && sIndex < sLength+1 {
				sIndex++
			}
			// If (string) matches (pattern),
			// when (string) matches (pattern)* as well
			dp[pIndex][sIndex-1] = dp[pIndex-1][sIndex-1]
			// If (string) matches (pattern), 
			// when (string)(whatever_characters) matches (pattern)* as well
			for sIndex < sLength+1 {
				dp[pIndex][sIndex] = true
				sIndex++
			}
		} else if p[pIndex-1] == '?' {
			// the current character in the pattern is '?'
			for sIndex = 1; sIndex < sLength+1; sIndex++ {
				dp[pIndex][sIndex] = dp[pIndex-1][sIndex-1]
			}
		} else {
			// the current character in the pattern is not '*' or '?'
			for sIndex = 1; sIndex < sLength+1; sIndex++ {
				// Match is possible if there is a previous match and current characters are the same
				dp[pIndex][sIndex] = dp[pIndex-1][sIndex-1] && p[pIndex-1] == s[sIndex-1]
			}
		}
	}
	return dp[pLength][sLength]
}
