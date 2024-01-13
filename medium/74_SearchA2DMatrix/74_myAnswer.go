package leetcode
/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。
该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
*/

/*
解法: 二分查找
把二维矩阵展开成一维数字，就是一个升序数组。


结果: 执行用时 :8 ms 内存消耗 :3.8 MB
*/

func searchMatrix(matrix [][]int, target int) bool {
	var row = len(matrix)
	if row == 0 {
		return false
	}
	var col = len(matrix[0])
	var total = row * col

	var left, right, mid = 0, total - 1, 0
	var i, j int
	for left <= right {
		mid = (left + right) / 2

		i = mid / col
		j = mid % col

		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
