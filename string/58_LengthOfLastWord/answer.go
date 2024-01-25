package LeetCode

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
*/

/*
解法: 倒序遍历，遇到字母时开始计数，之后遇到空格或遍历结束返回计数值


结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/

func lengthOfLastWord(s string) int {
	var ans int
	var length = len(s)
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' && ans == 0 {
			continue
		}
		if s[i] == ' ' {
			break
		}
		ans++
	}
	return ans
}
