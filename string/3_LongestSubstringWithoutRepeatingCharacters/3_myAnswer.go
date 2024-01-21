package LeetCode

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

import "strings"

/*
解法:
维护一个滑动窗口，窗口内的都是没有重复的字符，去尽可能的扩大窗口的大小，窗口不停的向右滑动。
（1）如果当前遍历到的字符从未出现过，那么直接扩大右边界；
（2）如果当前遍历到的字符出现过，则缩小窗口（左边索引向右移动），然后继续观察当前遍历到的字符；
（3）重复（1）（2），直到左边索引无法再移动；
（4）维护一个结果length，每次用出现过的窗口大小来更新结果length，最后返回length获取结果。

结果: 执行用时 :4 ms 内存消耗 :2.6 MB
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var length, left, right int
	var s1 string
	s1 = s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > length {
			length = len(s1)
		}
	}
	return length
}
