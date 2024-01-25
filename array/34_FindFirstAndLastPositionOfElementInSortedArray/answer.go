package LeetCode

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
*/

/*
解法: 二分查找


结果: 执行用时 :8 ms 内存消耗 :4.1 MB
*/

func searchRange(nums []int, target int) []int {
	var length = len(nums)
	if length == 0 {
		return []int{-1, -1}
	}

	var left, right, mid = 0, length - 1, (length - 1) >> 1
	for left <= right {
		mid = (left + right) >> 1
		if target == nums[mid] {
			var a, b = mid, mid
			for a > 0 {
				a--
				if target != nums[a] {
					a++
					break
				}
			}
			for b < length-1 {
				b++
				if target != nums[b] {
					b--
					break
				}
			}
			return []int{a, b}
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
