/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

进阶:如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
package leetcode

import "math"

/*
解法: 贪心


结果: 执行用时 :8 ms 内存消耗 :3.3 MB
*/


func maxSubArray(nums []int) int {
	var ans = math.MinInt32
	var length = len(nums)
	var sum int
	for i := 0; i < length; i++ {
		sum += nums[i]
		if ans < sum {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}