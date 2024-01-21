package LeetCode

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

进阶:
一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
*/

import "math"

/*
   解法: 暴力解，但不能直接改为0，否则会影响后续判断。
   遍历原始矩阵，如果发现如果某个元素 cell[i][j] 为 0，我们将第 i 行和第 j 列的所有非零元素设成很大的负虚拟值（比如说 -1000000）。注意，正确的虚拟值取值依赖于问题的约束，任何允许值范围外的数字都可以作为虚拟之。
   最后，我们遍历整个矩阵将所有等于虚拟值（常量在代码中初始化为 MODIFIED）的元素设为 0。


   结果: 执行用时 :16 ms 内存消耗 :6 MB
*/

func setZeroes(matrix [][]int) {
	// 标记值，纯粹是试出来的，这个值只要不存在原矩阵即可，但是如果要严谨地找起来很无语
	const zeroFlag = math.MinInt32 + 5

	var row = len(matrix)
	if row == 0 {
		return
	}
	var col = len(matrix[0])

	// 标记0值
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = zeroFlag
			}
		}
	}

	// 根据标记设置0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == zeroFlag {
				for k := 0; k < row; k++ {
					if matrix[k][j] != zeroFlag {
						matrix[k][j] = 0
					}
				}
				for k := 0; k < col; k++ {
					if matrix[i][k] != zeroFlag {
						matrix[i][k] = 0
					}
				}
				matrix[i][j] = 0
			}
		}
	}
}
