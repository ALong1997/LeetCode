/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
*/
package leetcode

/*
解法: 参考 T54 ，按层给矩阵赋值


结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/

func generateMatrix(n int) [][]int {
	var ans = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	var left, right, top, bottom, num = 0, n - 1, 0, n - 1, 1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			ans[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				ans[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				ans[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return ans
}