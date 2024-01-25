package LeetCode

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/

import "math"

/*
   解法: 不能使用除法，那就用减法代替吧
   先判断结果的正负号，将 dividend 和 divisor 都转化为正数更方便运算
   优化：每次减掉的值divisor翻倍，如果减超了就把divisor恢复原样重新慢慢翻倍


   结果: 执行用时 :0 ms 内存消耗 :2.5 MB
*/

func divide(dividend int, divisor int) int {
	var ans int
	var flag = true
	if dividend < 0 {
		dividend = -dividend
		flag = !flag
	}
	if divisor < 0 {
		divisor = -divisor
		flag = !flag
	}

	if divisor == 1 {
		ans = dividend
	} else {
		oldDivisor := divisor
		divisorCount := 1
		for dividend >= divisor {
			for dividend >= divisor {
				dividend -= divisor
				ans += divisorCount

				divisor += divisor
				divisorCount += divisorCount
			}
			divisor = oldDivisor
			divisorCount = 1
		}
	}
	if !flag {
		ans = -ans
		if ans < math.MinInt32 {
			return math.MinInt32
		}
	} else if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}
