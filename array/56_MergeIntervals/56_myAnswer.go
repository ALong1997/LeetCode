package leetcode
/*
给出一个区间的集合，请合并所有重叠的区间。
*/

import (
	"sort"
)

/*
   解法: 排序 暴力解


   结果: 执行用时 :12 ms 内存消耗 :4.7 MB
*/

func merge(intervals [][]int) [][]int {
	var ans [][]int
	var length = len(intervals)
	if length == 0 {
		return ans
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	var count int // len(ans)-1
	for i := 1; i < length; i++ {
		if intervals[i][0] > ans[count][1] {
			ans = append(ans, intervals[i])
			count++
		} else if intervals[i][1] > ans[count][1] {
			ans[count][1] = intervals[i][1]
		}
	}
	return ans
}
