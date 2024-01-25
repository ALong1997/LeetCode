package LeetCode

/*
给出一个无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/

/*
解法: 暴力解
将 intervals 分为三个部分，left，mid，right
left 在 newInterval 左，right 在 newInterval 右，mid 即 left+1 到 right-1 这一段区间
1.left + 1 = right 不需要合并区间直接插入，找到插入点插入即可
2.left + 1 < right 需要合并区间,mid 若存在则合并 mid 与 newInterval
注意考虑边界情况


结果: 执行用时 :12 ms 内存消耗 :5.9 MB
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}

	var ans [][]int
	var length = len(intervals)
	if length == 0 {
		return append(intervals, newInterval)
	}
	var left, right = -1, length
	for i := 0; i < length; i++ {
		if newInterval[0] > intervals[i][1] {
			left = i
			var temp = make([][]int, left+1)
			copy(temp, intervals[:left+1])
			ans = temp
			continue
		}
		if newInterval[1] < intervals[i][0] {
			right = i
			break
		}
	}

	if right != left+1 {
		if newInterval[0] > intervals[left+1][0] {
			newInterval[0] = intervals[left+1][0]
		}
		if newInterval[1] < intervals[right-1][1] {
			newInterval[1] = intervals[right-1][1]
		}
	}
	ans = append(ans, newInterval)
	if right != length {
		ans = append(ans, intervals[right:]...)
	}

	return ans
}
