package leetcode

/*
解法: 动态规划
参考 https://www.jianshu.com/p/2eeb6599ec30


结果: 执行用时 :0 ms 内存消耗 :2.4 MB
*/


func reference3IsMatch(s string, p string) bool {
	var col, row = len(s), len(p)
	//建立一个boolean数组，数组长度为 row + 1 * col + 1，多的第一列第一行代表空串的情况
	var arr = make([][]bool, row+1)
	for i := 0; i < row+1; i++ {
		arr[i] = make([]bool, col+1)
	}

	// 初始化第一行，第一行剩余元素全部变为false，因为pattern如果是空串，那么只要s不是空串，都不能匹配
	arr[0][0] = true
	//for i := 1; i < col+1; i++ {
	//	arr[0][i] = false
	//}
	//初始化第0列，此时s是空串，所以只能是x*y*这种形式的
	//注意，i-1代表的真实的p的索引
	for i := 1; i < row+1; i++ {
		arr[i][0] = i > 1 && p[i-1] == '*' && arr[i-2][0]
	}

	for i := 1; i < col+1; i++ {
		for j := 1; j < row+1; j++ {
			// 如果当前的pattern字符串是*，需要同时满足下面的两个条件，才能为true，此时s其实是固定住的，在思考的时候可以这么想
			//		1.如果往前数两行的位置可以匹配，那么这里的*可以代表空，此时可以匹配，
			//		2.如果s是aa，p是a*的话，我们不能根据往前数两行来判断，这时必须满足p[j-2]==s[i-1] 或者说 p[j-2]是'.'
			if p[j-1] == '*' {
				arr[j][i] = arr[j-2][i] || (p[j-2] == s[i-1] || p[j-2] == '.') && arr[j][i-1]
			} else {
				arr[j][i] = (p[j-1] == '.' || p[j-1] == s[i-1]) && arr[j-1][i-1]
			}
		}
	}
	return arr[row][col]
}