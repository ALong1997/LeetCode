/*
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/
package leetcode

/*
解法: 动态规划 - T84的优化
我们可以以常数时间计算出在给定的坐标结束的矩形的最大宽度。我们可以通过记录每一行中每一个方块连续的“1”的数量来实现这一点。
每遍历完一行，就更新该点的最大可能宽度。通过以下代码即可实现。 row[i] = row[i - 1] + 1 if row[i] == '1'.

一旦我们知道了每个点对应的最大宽度，我们就可以在线性时间内计算出以该点为右下角的最大矩形。
当我们遍历列时，可知从初始点到当前点矩形的最大宽度，就是我们遇到的每个最大宽度的最小值。
我们定义：
	maxWidth = min(maxWidth, widthHere)
	curArea = maxWidth * (currentRow - originalRow + 1)
	maxArea = max(maxArea, curArea)


结果: 执行用时 :12 ms 内存消耗 :4 MB
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var result int
	var dp = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				// compute the maximum width and update dp with it
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
				// compute the maximum area rectangle with a lower right corner at [i, j]
				var width = dp[i][j]
				for k := i; k >= 0; k-- {
					if width > dp[k][j] {
						width = dp[k][j]
					}
					if result < width*(i-k+1) {
						result = width * (i - k + 1)
					}
				}
			}
		}
	}

	return result
}
