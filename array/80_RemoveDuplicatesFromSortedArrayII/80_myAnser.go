/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/

/*
解法: 暴力解
需要注意边界处理（切片尾部）


结果: 执行用时 :4 ms 内存消耗 :3 MB
*/

func removeDuplicates(nums []int) int {
	if len(nums) > 2 {
		var count int
		for i := 0; i < len(nums)-2; i++ {
			if nums[i] == nums[i+1] && nums[i] == nums[i+2] {
				count++
				continue
			} else {
				i -= count
				nums = append(nums[:i], nums[i+count:]...)
				count = 0
			}
		}

		// 如果切片内最后几个数需要去重
		if count > 0 {
			nums = nums[:len(nums)-count]
		}
	}
	return len(nums)
}
