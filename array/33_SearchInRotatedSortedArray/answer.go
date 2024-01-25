package LeetCode

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。
*/

/*
解法: 二分查找，利用至少有半边数组保持有序的特性


结果: 执行用时 :4 ms 内存消耗 :2.6 MB
*/

func search(nums []int, target int) int {
	var length = len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	var left, right, mid = 0, length - 1, (length - 1) >> 1
	for left <= right {
		mid = (left + right) >> 1
		if target == nums[mid] {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
