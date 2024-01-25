package LeetCode

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
数独部分空格内已填入了数字，空白格用 '.' 表示。

说明:
一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。
*/

/*
解法: 一次遍历即可解决，需要用三个数组记录行、列、九宫格中的数字出现的次数
九宫格序号：indexBox := (i/3)*3+j/3


结果: 执行用时 :4 ms 内存消耗 :2.8 MB
*/

func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				indexBox := (i/3)*3 + j/3
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if sbox[indexBox][num] == 1 {
					return false
				} else {
					sbox[indexBox][num] = 1
				}
			}
		}
	}
	return true
}
