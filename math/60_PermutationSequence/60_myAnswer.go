package leetcode
/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：
给定 n 的范围是 [1, 9]
给定 k 的范围是 [1, n!]
*/

import "strconv"

/*
   解法: 找规律，根据 k 从最高位选择数字


   结果: 执行用时 :0 ms 内存消耗 :2 MB
*/

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i
		tokens[i-1] = i
	}

	k--
	var ans string
	for n > 0 {
		n--
		div := k / factorial[n]
		k = k % factorial[n]
		ans += strconv.Itoa(tokens[div])
		// 已使用的数字删掉
		tokens = append(tokens[:div], tokens[div+1:]...)
	}
	return ans
}
