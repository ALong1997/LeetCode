package LeetCode

import (
	"strconv"
)

/*
解法: 按计数分类
当且仅当它们的字符计数（每个字符的出现次数）相同时，两个字符串是字母异位词。


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

// 不用给字符串排序
// Using HashMap Using characters frequency as key
// Time Complexity: O(n*k) where k is the max length of string in strs
// Space Complexity: O(n*k)
func referenceGroupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 用一个大小为26的int数组来统计每个单词中字符出现的次数，
// 然后将int数组转为一个唯一的字符串，跟字符串数组进行映射，
func getKey(s string) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	key := ""
	for _, n := range freq {
		key += strconv.Itoa(n) + "/"
	}
	return key
}
