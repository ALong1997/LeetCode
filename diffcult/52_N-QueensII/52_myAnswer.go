/*
给定一个整数 n，返回 n 皇后不同的解决方案的数量。
*/
package leetcode

/*
解法: 参考T51


结果: 执行用时 :8 ms 内存消耗 :3.3 MB
*/


type queen struct {
	n int				// N
	stack map[int]int	// 某行皇后的纵坐标
	col []int			// 列是否被攻击
	pie map[int]int		// 副对角线是否被攻击
	na map[int]int		// 主对角线是否被攻击
	ret int				// 答案
}

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	q := NewQueen(n)
	q.backtrack(0)
	return q.ret
}

func NewQueen(n int) *queen {
	stack := make(map[int]int, n)
	col := make([]int, n)
	pie := make(map[int]int, n)
	na := make(map[int]int, n)
	return &queen{ n, stack, col ,pie, na, 0}
}

func (q *queen) backtrack(row int) {
	if q.n < 1 {
		return
	}
	for col :=0; col < q.n; col++ {
		if q.isNotUnderAttack(row, col) {
			q.addPos(row, col)
			if row + 1 == q.n {
				q.addSolution()
			} else {
				q.backtrack(row + 1)
			}
			q.rmPos(row, col)
		}
	}
}

// 是否不在皇后攻击范围内
func (q *queen) isNotUnderAttack(row, col int) bool {
	return q.col[col] + q.pie[row + col] + q.na[col - row + 2 * q.n ] == 0
}

func (q *queen) addPos(row, col int) {
	q.stack[row] = col
	q.col[col] = 1
	q.pie[row + col] = 1
	q.na[col - row + 2 * q.n] = 1
}

func (q *queen) rmPos(row, col int) {
	delete(q.stack, row)
	q.col[col] = 0
	q.pie[row + col] = 0
	q.na[col - row + 2 * q.n] = 0
}

func (q *queen) addSolution() {
	q.ret++
}
