package LeetCode

import (
	"fmt"
	"math/big"
)

/*
解法: 位运算
	把 a 和 b 转换成整型数字 x 和 y，在接下来的过程中，x 保存结果，y 保存进位。
	当进位不为 0 时
		计算当前 x 和 y 的无进位相加结果：answer = x ^ y
		计算当前 x 和 y 的进位：carry = (x & y) << 1
		完成本次循环，更新 x = answer，y = carry
	返回 x 的二进制形式


结果: 执行用时 :0 ms 内存消耗 :2.6 MB
*/

// 位运算 先用异或运算计算两数没有进位的结果，然后通过与运算和左移计算进位的结果
func referenceAddBinary(a string, b string) string {
	x, _ := new(big.Int).SetString(a, 2)
	y, _ := new(big.Int).SetString(b, 2)
	zero, _ := new(big.Int).SetString("0", 2)

	for y.Cmp(zero) != 0 {
		ans := new(big.Int).Xor(x, y)
		car := x.And(x, y).Lsh(x, 1)
		x, y = ans, car
	}
	return fmt.Sprintf("%b", x)
}
