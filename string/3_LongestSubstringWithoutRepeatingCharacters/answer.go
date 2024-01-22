package LeetCode

import "strings"

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
