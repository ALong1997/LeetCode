package LeetCode

/*
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。

下图是字符串 s1 = "great" 的一种可能的表示形式。
    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。

例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。

同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。

给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
*/

/*
解法: 动态规划 + 递归
直观的做法:  s是t的扰乱字符串等价于
	1: 左右节点不交换:  s[0:k]是t[0:k]的扰乱字符串 && s[k:]是t[k:]的扰乱字符串
	2: 左右节点交换:    s[0:n-k]是t[k:]的扰乱字符串 && s[n-k:]是t[0:k]的扰乱字符串
	其中k >= 1 && k <= length-1

结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	s1hash, s2hash := [26]int{}, [26]int{}
	for _, ch := range s1 {
		s1hash[ch-'a']++
	}
	for _, ch := range s2 {
		s2hash[ch-'a']++
	}
	for i := 0; i < 26; i++ {
		if s1hash[i] != s2hash[i] {
			return false
		}
	}

	n := len(s1)
	for length := 1; length <= n-1; length++ {
		if isScramble(s1[0:length], s2[0:length]) && isScramble(s1[length:], s2[length:]) {
			return true
		}

		if isScramble(s1[0:length], s2[n-length:]) && isScramble(s1[length:], s2[0:n-length]) {
			return true
		}
	}

	return false
}
