package LeetCode

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

说明：
所有输入均为小写字母。
不考虑答案输出的顺序。
*/

import (
	"sort"
)

/*
   解法: 排序数组分类
   当且仅当它们的排序字符串相等时，两个字符串是字母异位词。
   用一个map记录每组字母异位词中的字符串


   结果: 执行用时 :32 ms 内存消耗 :7.9 MB
*/

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	record := make(map[string][]string)
	for _, str := range strs {
		s := []rune(str)            // 把错位词的字符顺序重新排列，那么会得到相同的结果
		sort.Sort(sortRunes(s))     // 以此作为key，将所有错位词都保存到字符串数组中
		val, _ := record[string(s)] // 建立key和字符串数组之间的映射
		val = append(val, str)
		record[string(s)] = val
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
