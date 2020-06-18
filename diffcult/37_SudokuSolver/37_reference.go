/*
编写一个程序，通过已填充的空格来解决数独问题。
一个数独的解法需遵循如下规则：
	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	空白格用 '.' 表示。

说明:
给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
*/
package leetcode

/*
解法: 约束编程 + 回溯
约束编程基本的意思是在放置每个数字时都设置约束。在数独上放置一个数字后立即排除当前 行， 列 和 子方块 对该数字的使用。
这会传播 约束条件 并有利于减少需要考虑组合的个数。

九宫格序号：indexBox := (i/3)*3+j/3
回溯函数  backtrack(row = 0, col = 0)。
	从最左上角的方格开始 row = 0, col = 0。直到到达一个空方格。
	从1 到 9 迭代循环数组，尝试放置数字 d 进入 (row, col) 的格子。
		如果数字 d 还没有出现在当前行，列和子方块中：
			如果这是最后一个格子row >= 9, col >= 9 ：意味着已经找出了数独的解。
			否则放置接下来的数字。
			将 d 放入 (row, col) 格子中。记录下 d 已经出现在当前行，列和子方块中。
			如果数独的解还没找到,将最后的数从 (row, col) 移除。


结果: 执行用时 :4 ms 内存消耗 :2.1 MB
*/

type sudoMap struct {
	RowMap [9][9]int
	ColMap [9][9]int
	BoardMap [9][9]int
}

func solveSudoku(board [][]byte)  {
	// 初始化
	smap := new(sudoMap)

	// 将 board 构建至 smap
	for r, rowData := range board {
		for c, cell := range rowData {
			if cell == '.' {
				continue
			}
			intV := cell2Int(cell)
			putNumber(smap, intV, r, c)
		}
	}

	backtrace(board, smap, 0, 0)
}

func backtrace(board [][]byte, smap *sudoMap, r, c int) bool {
	if isEndPos(r, c) {
		return true
	}

	nr, nc := nextPos(r, c)
	if board[r][c] == '.' {
		for i := 1; i <= 9; i++ {
			if !canPutNumber(smap, i, r, c) {
				continue
			}

			// 设置值
			putNumber(smap, i, r, c)
			board[r][c] = byte(i + '0')

			// 下一个
			success := backtrace(board, smap, nr, nc)

			if !success {
				// 清除元素
				clearNumber(smap, i, r, c)
				board[r][c] = '.'
			}else {
				return true
			}
		}
		return false
	}else {
		return backtrace(board, smap, nr, nc)
	}
}

func canPutNumber(smap *sudoMap, v, r, c int) bool {
	boardPos := rc2Board(r, c)
	return smap.RowMap[r][v - 1] == 0 && smap.ColMap[c][v - 1] == 0 && smap.BoardMap[boardPos][v - 1] == 0
}

func putNumber(smap *sudoMap, v, r, c int) {
	boardPos := rc2Board(r, c)
	smap.RowMap[r][v - 1] += 1
	smap.ColMap[c][v - 1] += 1
	smap.BoardMap[boardPos][v - 1] += 1
}

func clearNumber(smap *sudoMap, v, r, c int) {
	boardPos := rc2Board(r, c)
	smap.RowMap[r][v - 1] = 0
	smap.ColMap[c][v - 1] = 0
	smap.BoardMap[boardPos][v - 1] = 0
}

func nextPos(r, c int) (int, int) {
	nr := r
	nc := c + 1
	if nc == 9 {
		nc = 0
		nr += 1
	}

	return nr, nc
}

func isEndPos(r, c int) bool {
	return r >= 9 || c >= 9
}

func cell2Int(b byte) int {
	return int(b - '0')
}

func rc2Board(r, c int) int {
	return (r / 3) * 3 + c / 3
}