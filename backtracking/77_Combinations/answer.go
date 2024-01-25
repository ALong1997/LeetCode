package LeetCode

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/

/*
解法: 回溯
这是一个回溯法函数，它将第一个添加到组合中的数和现有的组合作为参数。
backtrack(first, curr)
	若组合完成，添加到输出中。
	遍历从 first 到 n 的所有整数。
		将整数 i 添加到现有组合 curr 中。
		继续向组合中添加更多整数 : backtrack(i + 1, curr).
		将 i 从 curr 中移除，实现回溯。


结果: 执行用时 :8 ms 内存消耗 :6.1 MB
*/

var ans [][]int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	if n < k || n == 0 || k == 0 {
		return ans
	}

	getCombinations(n, k, 1, 0, nil)
	return ans
}

func getCombinations(n int, k int, start int, cnt int, t []int) {
	if cnt == k {
		s := make([]int, len(t))
		copy(s, t)
		ans = append(ans, s)
		return
	}

	// 剪枝："i<=n"  -->  "i<=n-(k-cnt)+1"
	for i := start; i <= n-(k-cnt)+1; i++ {
		t = append(t, i)
		getCombinations(n, k, i+1, cnt+1, t)
		t = t[:len(t)-1]
	}
}
