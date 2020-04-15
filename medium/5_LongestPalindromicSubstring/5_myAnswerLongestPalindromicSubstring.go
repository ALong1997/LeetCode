//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
package leetcode

/*
解法: 暴力解


结果: 执行用时 :1116 ms 内存消耗 :2.2 MB
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	longest := string(s[0])
	for i := 0; i < length - 1; i++ {
		for j := i + 1; j < length; j++ {
			if isPalindromic(s[i : j + 1]) && j - i + 1 > len(longest) {
				longest = s[i : j + 1]
			}
		}
	}
	return longest
}

func isPalindromic(s string) bool {
	length := len(s)
	for i := 0; i < length / 2; i++ {
		if s[i] != s[length - i - 1] {
			return false
		}
	}
	return true
}