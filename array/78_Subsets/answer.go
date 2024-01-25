package LeetCode

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。
*/

/*
解法: 回溯


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

func subsets(nums []int) [][]int {
	var ans [][]int

	var backtrace func(pos, num int, cur []int)
	backtrace = func(pos, num int, cur []int) {
		if len(cur) == num {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := pos; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrace(i+1, num, cur)
			cur = cur[:len(cur)-1]
		}
	}

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		backtrace(0, i, cur)
	}

	return ans
}
