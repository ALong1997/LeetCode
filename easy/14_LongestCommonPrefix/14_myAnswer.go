/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
package leetcode

/*
解法: 暴力解


结果: 执行用时 :0 ms 内存消耗 :2.4 MB
*/

func longestCommonPrefix(strs []string) string {
	var s string
	if len(strs) == 0 {
		return s
	} else if len(strs) == 1 {
		s = strs[0]
	} else {
		temp := strs[0]
		for i := 0; i < len(temp); i++ {
			c := temp[i]
			for j := 1; j < len(strs) ; j++ {
				if i >= len(strs[j]) || strs[j][i] != c {
					return s
				}
			}
			s += string(c)
		}
	}
	return s
}