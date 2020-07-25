/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/
package leetcode

/*
解法:参考T62
dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]


结果: 执行用时 :16 ms 内存消耗 :4.8 MB
*/

func minPathSum(grid [][]int) int {
	var ans int

	var m = len(grid)
	if m == 0 {
		return ans
	}

	var n = len(grid[0])
	cur := make([]int,n)

	cur[0] = grid[0][0]
	for i:=1;i<n;i++{
		cur[i] = grid[0][i] + cur[i-1]
	}

	for i:=1;i<m;i++{
		cur[0] += grid[i][0]
		for j:=1;j<n;j++{
			if cur[j] < cur[j-1] {
				cur[j] = cur[j] + grid[i][j]
			} else {
				cur[j] = cur[j-1] + grid[i][j]
			}
		}
	}
	return cur[n-1]
}