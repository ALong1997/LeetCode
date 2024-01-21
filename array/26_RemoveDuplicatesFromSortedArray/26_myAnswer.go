package LeetCode

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

/*
解法: 暴力解
为了避免消耗时间过多，没有采用 append函数删除数组元素的做法


结果: 执行用时 :0 ms 内存消耗 :4.6 MB
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[ans] != nums[i] {
			ans++
			nums[ans] = nums[i]
		}
	}
	return ans + 1
}
