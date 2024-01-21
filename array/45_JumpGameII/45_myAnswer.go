package LeetCode

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

/*
解法: 动态规划
dp[i]记录 任一点 到 i 的最短路径，0 表示不是最短路径
从 i 走到 j 时
		dp[i] + 1,				dp[j]==0
dp[j] =
		min(dp[i] + 1, dp[j]),	dp[j]!=0

结果: 执行用时 :428 ms 内存消耗 :4.7 MB
*/

func jump(nums []int) int {
	var step int
	var length = len(nums)
	var dp = make([]int, length)

	for i := 0; i < length; i++ {
		for j := 1; j <= nums[i] && i+j < length; j++ {
			step = dp[i] + 1
			if dp[i+j] == 0 || step < dp[i+j] {
				dp[i+j] = step
			}
		}
	}
	return dp[length-1]
}
