/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
注意:假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0
*/
package leetcode

import "math"

/*
解法: 暴力解


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if y < math.MinInt32 || y > math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	return y
}