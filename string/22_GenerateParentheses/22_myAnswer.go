package leetcode
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

/*
解法: 回溯法
括号成对存在，先有左括号再有右括号，所以只有右括号的数量小于左括号才进行右括号的添加。
最后如果右括号的数量等于0，表示右括号已经排完了，同时意味着左括号也排完了。


结果: 执行用时 :0 ms 内存消耗 :2.8 MB
*/

func generateParenthesis(n int) []string {
	// 使用new方法开辟内存返回内存地址
	res := new([]string)

	backtracking(n, n, "", res)

	return *res
}

func backtracking(left, right int, tmp string, res *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		*res = append(*res, tmp)
		return
	}

	// 生成左括号
	if left > 0 {
		backtracking(left-1, right, tmp+"(", res)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backtracking(left, right-1, tmp+")", res)
	}
}
