package leetcode

import "sort"

/*
解法: 排序 + 双指针
1.特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回。
2.对数组进行排序。
3.遍历排序后数组：
	若 nums[i]>0 :因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
	对于重复元素：跳过，避免出现重复解
	令左指针 L=i+1，右指针 R=n−1，当 L<R 时，执行循环：
		当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
		若和大于 0，说明 nums[R] 太大，R 左移
		若和小于 0，说明 nums[L] 太小，L 右移


结果: 执行用时 :36 ms 内存消耗 :7 MB
*/

func referenceThreeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) > 2 {
		// 排序
		sort.Ints(nums)
		var left, right, temp int
		for i := 0; i < len(nums); i++ {
			if nums[i] > 0 {
				return ans
			}
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			left = i + 1
			right = len(nums) - 1
			for left < right {
				temp = nums[left] + nums[i] + nums[right]
				if temp == 0 {
					ans = append(ans, []int{nums[left], nums[i], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if temp < 0 {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}