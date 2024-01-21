package leetcode

/*
解法: 位运算将二维数组简化为一维，降低空间复杂度
本题可以使用一个9位二进制数判断数字是否被访问。第 k 位数为 1 代表已加入，为 0 代表未加入
更新方式(记九位数为 val，传入的数字为 n
	判断是否加入：将九位数右移位 n 位，与 1 进行与运算
		结果为0：未加入，将传入的数字加入九位数
		结果为1：已加入，返回false
	将传入的数字加入九位数：将 1 左移位 n 位，与 val 异或即可


结果: 执行用时 :4 ms 内存消耗 :2.8 MB
*/

func referenceIsValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3  // 3x3的块索引号
			if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}