/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。
*/
package leetcode

import (
	"sort"
)

/*
解法: 类似T15
(1) 第一层循环选择第一个数固定然后去遍历其它三个数。
(2) 第二层循环选择第二个数固定然后去遍历最后两个数。
(3) 通过双指针选出最后符合条件的两个数。

结果: 执行用时 :4 ms 内存消耗 :2.7 MB
*/
func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	size := len(nums)
	if size > 3 {
		// 排序
		sort.Ints(nums)
		if target < nums[0]+nums[1]+nums[2]+nums[3] || target > nums[size-1]+nums[size-2]+nums[size-3]+nums[size-4] {
			return ans
		}
		var left, temp, right int
		for i := 0; i < size - 3; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			if target < nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] {
				break
			}
			for j := i + 1; j < size - 2; j++ {
				if j > i + 1 && nums[j] == nums[j - 1] {
					continue
				}
				if target < nums[i] + nums[j] + nums[j + 1] + nums[j + 2] {
					break
				}
				if target > nums[i] + nums[j] + nums[size - 1] + nums[size - 2] {
					continue
				}
				left = j + 1
				right = size - 1
				for left < right {
					temp = nums[left] + nums[i] + nums[j] + nums[right]
					if temp == target {
						ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
						for left < right && nums[left] == nums[left+1] {
							left++
						}
						for left < right && nums[right] == nums[right-1] {
							right--
						}
						left++
						right--
					} else if temp < target {
						left++
					} else {
						right--
					}
				}
			}
		}
	}
	return ans
}