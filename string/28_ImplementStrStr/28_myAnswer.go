package LeetCode

/*
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

/*
解法: KMP
本题解法很多，也有更方便的解法，不过我想熟悉下KMP算法


结果: 执行用时 :0 ms 内存消耗 :2.6 MB
*/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	next := getNextArr(needle)
	i, j := 0, 0
	for ; i < len(haystack); i++ {
		for haystack[i] != needle[j] && j > 0 {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - len(needle) + 1
			}
		}
	}
	return -1
}

func getNextArr(s string) []int {
	res := make([]int, len(s))
	maxMatchLen := 0
	for i := 1; i < len(s); i++ {
		for maxMatchLen > 0 && s[maxMatchLen] != s[i] {
			maxMatchLen = res[maxMatchLen-1]
		}

		if s[maxMatchLen] == s[i] {
			maxMatchLen++
		}

		res[i] = maxMatchLen
	}
	return res
}
