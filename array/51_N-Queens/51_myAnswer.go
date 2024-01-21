package leetcode
/*
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

提示：皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。
*/

/*
解法: 约束编程 + 回溯
约束编程:在放置每个皇后以后增加限制。当在棋盘上放置了一个皇后后，立即排除当前行，列和对应的两个对角线。该过程传递了 约束 从而有助于减少需要考虑情况数。

在建立算法之前，我们来考虑两个有用的细节。
	一行只可能有一个皇后且一列也只可能有一个皇后。这意味着没有必要再棋盘上考虑所有的方格。只需要按列循环即可。
	对于所有的主对角线有 行号 - 列号 = 常数，对于所有的次对角线有 行号 + 列号 = 常数。这可以让我们标记已经在攻击范围下的对角线并且检查一个方格 (行号, 列号) 是否处在攻击位置。

现在已经可以写回溯函数 backtrack(row = 0).
	从第一个 row = 0 开始.
	循环列并且试图在每个 column 中放置皇后.
		如果方格 (row, column) 不在攻击范围内，在 (row, column) 方格上放置皇后。排除对应行，列和两个对角线的位置。
		If 所有的行被考虑过，row == N 意味着我们找到了一个解
			Else 继续考虑接下来的皇后放置 backtrack(row + 1).
		回溯：将在 (row, column) 方格的皇后移除.


结果: 执行用时 :8 ms 内存消耗 :3.3 MB
*/

type queen struct {
	n     int         // N
	stack map[int]int // 某行皇后的纵坐标
	col   []int       // 列是否被攻击
	pie   map[int]int // 副对角线是否被攻击
	na    map[int]int // 主对角线是否被攻击
	ret   [][]string  // 答案
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	q := newQueen(n)
	q.backtrack(0)
	return q.ret
}

func newQueen(n int) *queen {
	stack := make(map[int]int, n)
	col := make([]int, n)
	pie := make(map[int]int, n)
	na := make(map[int]int, n)
	return &queen{n, stack, col, pie, na, [][]string{}}
}

func (q *queen) backtrack(row int) {
	if q.n < 1 {
		return
	}
	for col := 0; col < q.n; col++ {
		if q.isNotUnderAttack(row, col) {
			q.addPos(row, col)
			if row+1 == q.n {
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
	return q.col[col]+q.pie[row+col]+q.na[col-row+2*q.n] == 0
}

func (q *queen) addPos(row, col int) {
	q.stack[row] = col
	q.col[col] = 1
	q.pie[row+col] = 1
	q.na[col-row+2*q.n] = 1
}

func (q *queen) rmPos(row, col int) {
	delete(q.stack, row)
	q.col[col] = 0
	q.pie[row+col] = 0
	q.na[col-row+2*q.n] = 0
}

func (q *queen) addSolution() {
	solution := make([]string, q.n)
	for row := 0; row < q.n; row++ {
		rowCol := q.stack[row]
		output := ""
		for col := 0; col < q.n; col++ {
			if rowCol == col {
				output += "Q"
			} else {
				output += "."
			}
		}
		solution[row] = output
	}
	q.ret = append(q.ret, solution)
}
