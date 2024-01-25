package LeetCode

/*
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。

说明：
你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
*/

/*
解法: 先转置，再翻转每一行

从行的角度看：顺时针旋转 90 度
原矩阵的第一行，变成了目标矩阵的最后一列。
原矩阵的第二行，变成了目标矩阵的倒数第二列。
原矩阵的第三行，变成了目标矩阵的倒数第三列。
原矩阵的第四行，变成了目标矩阵的倒数第四列。

矩阵顺时针旋转90的结果，就是把行变成列，把列变成行（即转置）。然后翻转对应的行即可。


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

func rotate(matrix [][]int) {
	var col = len(matrix[0])
	for i := 0; i < col; i++ {
		for j := i; j < col; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < col; i++ {
		for j := 0; j < col/2; j++ {
			matrix[i][j], matrix[i][col-j-1] = matrix[i][col-j-1], matrix[i][j]
		}
	}
}
