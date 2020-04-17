package leetcode

/*
解法: 动态规划
对于母串 s , 用 c[i,j] = 1 表示字串 s[i:j] 为回文串，则有递推公式
          c[i+1,j-1], if s[i] == s[j]
c[i,j] =
          0         , if s[i] != s[j]

特别地，对于这样的字符串——只包含单个字符、或者两个字符重复，其均为回文串：
c[i,i] = 1
c[i,i+1] = 1, if s[i] == s[i+1]


结果: 执行用时 :44 ms 内存消耗 :6.5 MB
*/
func referenceLongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	longest := string(s[0])
	c := make([][]bool, length)
	for i := range c {
		c[i] = make([]bool, length)
	}
	for gap := 0; gap < length ; gap++ {
		for i, j := 0, 0; i < length - gap; i++ {
			j = i + gap
			if s[i] == s[j] && (j - i <= 2 || c[i+1][j-1]) {
				c[i][j] = true
				if j - i + 1 > len(longest) {
					longest = s[i:j + 1]
				}
			}
		}
	}
	return longest
}


/*
解法: 分治法
回文串是左右对称的，如果从中心轴开始遍历，会减少一层循环。
思路：依次以母串的每一个字符为中心轴，得到回文串；然后通过比较得到最长的那一个。
注意：要考虑到像bacccccccab这样的特殊子串，可以理解它的中心轴是ccccccc。


结果: 执行用时 :4 ms 内存消耗 :2.2 MB
*/

func reference2LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	longest := string(s[0])
	for i := 0; i < length - 1; i++  {
		subS := palindrome(s, i)
		if len(subS) > len(longest) {
			longest = subS
		}
	}
	return longest
}

func palindrome(s string, mid int) string {
	l := mid - 1; r := mid + 1
	length := len(s)
	for r < length && s[r] == s[mid] {
		r++
	}
	for l >= 0 && r < length && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1:r]
}