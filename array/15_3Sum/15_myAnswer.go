package LeetCode

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

import (
	"sort"
)

/*
解法: (本方法几乎超时)
a + b + c = 0 <=> c = -(a + b) 即判断 -(a + b) 是否存在数组中
先对数组排序, 排序后若全为正数或全为负数则直接返回空。
由于数组已有序, 再对数组进行遍历查找即可。
查找出符合条件的结果，去重即可。

结果: 执行用时 :2628 ms 内存消耗 :6.9 MB
*/
func threeSum(nums []int) [][]int {
	var ans [][]int
	var a, b, c int
	if len(nums) > 2 {
		// 排序
		sort.Ints(nums)
		if nums[0] > 0 || nums[len(nums)-1] < 0 {
			return ans
		}
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				a = nums[i]
				b = nums[j]
				c = -(a + b)
				for k := j + 1; k < len(nums); k++ {
					if c == nums[k] {
						if len(ans) == 0 {
							ans = append(ans, []int{a, b, c})
						} else {
							flag := true
							for _, v := range ans {
								if v[0] == a && v[1] == b && v[2] == c {
									flag = false
									break
								}
							}
							if flag {
								ans = append(ans, []int{a, b, c})
							}
						}
						break
					}
				}
			}
		}
	}
	return ans
}
