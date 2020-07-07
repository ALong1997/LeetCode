/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。
*/
package leetcode

/*
解法: 参考T62
动态规划


结果: 执行用时 :0 ms 内存消耗 :2.5 MB
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m = len(obstacleGrid)
	if m == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	var n = len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一行
	for i := 0; i < n; i++ {
		dp[0][i] = 1
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
			for ; i < n; i++ {
				dp[0][i] = 0
			}
		}
	}

	// 初始化第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			for ; i < m; i++ {
				dp[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}