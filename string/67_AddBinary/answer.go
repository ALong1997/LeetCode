package LeetCode

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 1 和 0。

提示：
	每个字符串仅由字符 '0' 或 '1' 组成。
	1 <= a.length, b.length <= 10^4
	字符串如果不是 "0" ，就都不含前导零。
*/

/*
解法: 暴力解


结果: 执行用时 :0 ms 内存消耗 :2.6 MB
*/

func addBinary(a string, b string) string {
	var result = ""
	var flag = 0 // 存储进位
	var i, j = len(a) - 1, len(b) - 1

	for i >= 0 || j >= 0 {
		var t1, t2 = 0, 0
		if i >= 0 {
			t1 = int(a[i] - '0')
		}
		if j >= 0 {
			t2 = int(b[j] - '0')
		}
		sum := t1 + t2 + flag // 计算当前位置
		switch sum {
		case 3:
			flag = 1
			result = "1" + result
		case 2:
			flag = 1
			result = "0" + result
		case 1:
			flag = 0
			result = "1" + result
		case 0:
			flag = 0
			result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 { // 最终进位
		result = "1" + result
	}
	return result
}
