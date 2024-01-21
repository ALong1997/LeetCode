package leetcode

/*
解法: 字典序
将每个子集映射到长度为 n 的位掩码中，其中第 i 位掩码 nums[i] 为 1，表示第 i 个元素在子集中；如果第 i 位掩码 nums[i] 为 0，表示第 i 个元素不在子集中。
例如，位掩码 0..00（全 0）表示空子集，位掩码 1..11（全 1）表示输入数组 nums。
因此要生成所有子集，只需要生成从 0..00 到 1..11 的所有 n 位掩码。

算法
	生成所有长度为 n 的二进制位掩码。
	将每个子集都映射到一个位掩码数：位掩码中第 i 位如果是 1 表示子集中存在 nums[i]，0 表示子集中不存在 nums[i]。
	返回子集列表。


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

func referenceSubsets(nums []int) [][]int {
	var ans [][]int
	var length, max = len(nums), 1
	for i := 0; i < length; i++ {
		max = max << 1
	}
	for idx := 0; idx < max; idx++ {
		var col []int
		var posIdx = idx
		pos := 0
		for posIdx > 0 {
			if posIdx&1 == 1 {
				col = append(col, nums[pos])
			}
			posIdx = posIdx >> 1
			pos++
		}
		ans = append(ans, col)
	}
	return ans
}
