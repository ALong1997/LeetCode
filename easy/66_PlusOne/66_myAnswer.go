/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
package leetcode

/*
解法: 暴力解
倒序+1，注意进位即可。


结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/

func plusOne(digits []int) []int {
	var length = len(digits)
	var carry bool

	for i := length-1; i >= 0; i-- {
		if digits[i] == 9 {
			carry = true
		} else {
			carry = false
		}

		if carry {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	if carry {
		digits = append([]int{1}, digits...)
	}

	return digits
}