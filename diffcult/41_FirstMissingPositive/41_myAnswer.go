/*
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

注意：你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
*/
package leetcode

import "sort"

/*
解法: 排序后查找


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/


func firstMissingPositive(nums []int) int {
	var ans = 1
	sort.Ints(nums)
	for _, v := range nums {
		if v == ans {
			ans++
		}
	}
	return ans
}
