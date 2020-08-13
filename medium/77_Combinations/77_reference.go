package leetcode

/*
解法: 字典序
原理就是把数字组合当做一个整数，从小到大排序，如12、13、23、14、24、34。实现时首先找到最小的最高位数字，然后依次找次低位的最小数字，依次类推
主要思路是以字典序的顺序获得全部组合。
算法非常直截了当 :
	将 nums 初始化为从 1 到 k 的整数序列。 将 n + 1 添加为末尾元素，起到“哨兵”的作用。
	将指针设为列表的开头 j = 0.
	While j < k :
		将nums 中的前k个元素添加到输出中，换而言之，除了“哨兵”之外的全部元素。
		找到nums中的第一个满足 nums[j] + 1 != nums[j + 1]的元素，并将其加一，nums[j]++ 以转到下一个组合。


结果: 执行用时 :8 ms 内存消耗 :6.3 MB
*/

func referenceCombine(n int, k int) [][]int {
	var nums = make([]int, k+1)
	for i := 0; i < k+1; i++ {
		nums[i] = i + 1
	}
	nums[k] = n + 1
	var ans [][]int
	var j int
	for j < k {
		s := make([]int, k)
		copy(s, nums[:k])
		ans = append(ans, s)
		j = 0
		for j < k && nums[j+1] == nums[j]+1 {
			nums[j] = j + 1
			j++
		}
		nums[j]++
	}
	return ans
}
