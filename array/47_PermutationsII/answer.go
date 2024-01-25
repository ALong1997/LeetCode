package LeetCode

/*
给定一个 可包含重复数字 的序列，返回所有不重复的全排列。
*/

/*
解法: 回溯+剪枝 参考T46
什么情况需要剪枝？画树状图分析发现，同一个父节点下的所有子树重复数字之间交换位置才会产生重复结果。所以每次递归都需要维护一个 map 记录该数字是否使用过。
在回溯法求解全排列的基础上，添加了一个 map 重复，回溯法经典的思想，每次递归中维护一个 map 来检查是否在此阶段将两个相同的数交换了位置，若交换了则必定有所重复。


结果: 执行用时 :0 ms 内存消耗 :3.5 MB
*/

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var length = len(nums)

	var backtrack func(int)
	backtrack = func(first int) {
		if first == length {
			var temp = make([]int, length)
			copy(temp, nums)
			ans = append(ans, temp)
		}

		// 记录数字是否使用过的map
		var haveUse = make(map[int]bool, length)
		for i := first; i < length; i++ {
			if _, ok := haveUse[nums[i]]; !ok {
				haveUse[nums[i]] = true
				// 动态维护数组
				nums[first], nums[i] = nums[i], nums[first]
				// 继续递归填下一个数
				backtrack(first + 1)
				// 撤销操作
				nums[first], nums[i] = nums[i], nums[first]
			}
		}
	}
	backtrack(0)
	return ans
}
