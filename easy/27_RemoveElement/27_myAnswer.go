/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
package leetcode

/*
解法: 暴力解
为了避免消耗时间过多，没有采用 append函数删除数组元素的做法


结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/


func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}