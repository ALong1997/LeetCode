/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。
假定每组输入只存在唯一答案。
*/
package leetcode

import (
	"math"
	"sort"
)

/*
解法: 和T15类似的思路，排序 + 双指针


结果: 执行用时 :4 ms 内存消耗 :2.7 MB
*/
func threeSumClosest(nums []int, target int) int {
	var ans int
	if len(nums) > 2 {
		// 排序
		sort.Ints(nums)
		ans = nums[0] + nums[1] + nums[2]
		var left, right, temp int
		for i := 0; i < len(nums); i++ {
			left = i + 1
			right = len(nums) - 1
			for left < right {
				temp = nums[left] + nums[i] + nums[right]
				if math.Abs(float64(temp - target)) < math.Abs(float64(ans - target)) {
					ans = temp
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
				}
				if temp < target {
					left++
				} else if temp > target {
					right--
				} else {
					return target
				}
			}
		}
	}
	return ans
}