/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
行数为 4 时，排列如下：
L     D     R
E   O E   I I
E C   I H   N
T     S     G
输出: "LDREOEIIECIHNTSG"
*/
package leetcode

import (
	"bytes"
)

/*
解法: 寻找输入与输出两个字符串数组之间的对应规律，按行访问
第一行和最后一行每隔 (numRows - 1) << 1 个字符会有一个字符加入结果字符串
其他行每隔 (numRows - 1) << 1 个字符有两个字符加入结果字符串


结果: 执行用时 :8 ms 内存消耗 :4.1 MB
*/
func convert(s string, numRows int) string {
	length := len(s)
	if length <= numRows || numRows == 1 {
		return s
	}
	var buffer bytes.Buffer
	gap := (numRows - 1) << 1
	for i := 0; i < numRows; i++ {
		for j := i; j < length; j += gap {
			// 第一行和最后一行特殊处理
			if i == 0 || i == numRows - 1 {
				buffer.WriteByte(s[j])
			} else {
				buffer.WriteByte(s[j])
				if j + gap - i << 1 < length {
					buffer.WriteByte(s[j + gap - i << 1])
				}
			}
		}
	}
	return buffer.String()
}