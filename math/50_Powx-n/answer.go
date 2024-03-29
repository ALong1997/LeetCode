package LeetCode

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

说明:
-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

/*
解法: 分治 快速幂 + 递归
举个例子，如果我们要计算 x^64,从 x 开始，每次直接把上一次的结果进行平方，计算 6 次就可以得到 x^64 的值
再举一个例子，如果我们要计算 x^77, x^2 → x^4 → x^9 → x^19 → x^38 → x^77,这些步骤中，我们把上一次的结果进行平方后，还要额外乘一个 x

直接从左到右进行推导看上去很困难，因为在每一步中，我们不知道在将上一次的结果平方之后，还需不需要额外乘 x。但如果我们从右往左看，分治的思想就十分明显了：
当我们要计算 x^n 时，我们可以先递归地计算出 y = x^(⌊ n/2 ⌋)
根据递归计算的结果，如果 n 为偶数，那么 x^n = y^2; 如果 n 为奇数，那么 x^n = y^2 * x
递归的边界为 n=0，任意数的 0 次方均为 1。
由于每次递归都会使得指数减少一半，因此递归的层数为 O(log n)，算法可以在很快的时间内得到结果。


结果: 执行用时 :0 ms 内存消耗 :2 MB
*/

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
