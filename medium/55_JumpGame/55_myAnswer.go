package leetcode
/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
*/

/*
解法: 回溯 + 剪枝


结果: 执行用时 :24 ms 内存消耗 :8.5 MB
*/

func canJump(nums []int) bool {
	return backtrack(nums, len(nums), 0)
}

func backtrack(nums []int, length, index int) bool {
	if index >= length-1 {
		return true
	}
	var newIndex, count int
	for i := nums[index]; i > 0; i-- {
		newIndex = index + i
		if newIndex >= length-1 {
			return true
		}
		count++ // 剪枝
		if count <= nums[newIndex] && backtrack(nums, length, newIndex) {
			return true
		}
	}
	return false
}
