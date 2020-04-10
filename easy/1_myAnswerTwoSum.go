/*
给定一个整数数组 nums 和一个目标值 target
请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/
package leetcode

import "sort"

/*
解法：首尾递进查找
结果:
执行用时 :4 ms
内存消耗 :3.1 MB
*/
func twoSum(nums []int, target int) []int {
	n := len(nums)
	numsSort := make([]int, n)
	copy(numsSort, nums)
	sort.Ints(numsSort)

	for i, j := 0, n - 1; i < j; {
		sum := numsSort[i] +numsSort[j]
		if sum == target {
			a := find(nums, numsSort[i], -1)
			b := find(nums, numsSort[j], a)
			return []int{a, b}
		}
		if sum < target {
			i++
		}
		if sum > target {
			j--
		}
	}
	return []int{}
}

func find(nums []int, target, have int) int {
	for k, v := range nums {
		if v == target && k != have {
			return k
		}
	}
	return 0
}