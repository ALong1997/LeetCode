package LeetCode

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	longest := string(s[0])
	for i := 0; i < length-1; i++ {
		subS := palindrome(s, i)
		if len(subS) > len(longest) {
			longest = subS
		}
	}
	return longest
}

func palindrome(s string, mid int) string {
	l := mid - 1
	r := mid + 1
	length := len(s)
	for r < length && s[r] == s[mid] {
		r++
	}
	for l >= 0 && r < length && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
