package leetcode
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
*/

/*
解法: 二分查找
注意插入位置的边界，将 left == right 的情况单独处理


结果: 执行用时 :0 ms 内存消耗 :3.1 MB
*/

func searchInsert(nums []int, target int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}

	var left, right, mid = 0, length - 1, (length - 1) >> 1
	for left < right {
		mid = (left + right) >> 1
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 此时left = right
	if target <= nums[left] {
		return left
	}
	return left + 1
}
