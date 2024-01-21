package leetcode
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 10 ^ 9
*/

/*
解法: 排列组合
因为机器到底右下角，向下几步，向右几步都是固定的，
比如，m=3, n=2，我们只要向下 1 步，向右 2 步就一定能到达终点。
所以答案是组合数 C[m-1,m+n-2]


结果: 执行用时 :0 ms 内存消耗 :2 MB
*/

func uniquePaths(m int, n int) int {
	var ans int
	if n < 1 {
		return ans
	}

	var M, N = m - 1, m + n - 2
	if M<<1 > N {
		var a, b = 1, 1
		for i := N; i > M; i-- {
			a *= i
		}
		for i := N - M; i > 0; i-- {
			b *= i
		}
		ans = a / b
	} else {
		var a, b = 1, 1
		for i := N; i > N-M; i-- {
			a *= i
		}
		for i := M; i > 0; i-- {
			b *= i
		}
		ans = a / b
	}

	return ans
}
