package leetcode
/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
*/

/*
解法: 参考T33
本题是T33的升级版，本题中 nums 内可能出现重复元素


结果: 执行用时 :4 ms 内存消耗 :3.2 MB
*/

func search(nums []int, target int) bool {
	var length = len(nums)
	if length == 0 {
		return false
	}
	if length == 1 {
		if target == nums[0] {
			return true
		} else {
			return false
		}
	}

	var left, right, mid = 0, length - 1, (length - 1) >> 1
	for left <= right {
		// 去左右边界重复值
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		mid = (left + right) >> 1
		if target == nums[mid] {
			return true
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

	return false
}
