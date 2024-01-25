package LeetCode

/*
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

注意：整数序列中的每一项将表示为一个字符串。
*/

import (
	"bytes"
	"strconv"
)

/*
   解法: 递归


   结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		last := countAndSay(n - 1)
		first := last[0]
		num := 1
		var buffer bytes.Buffer
		for i := 1; i < len(last); i++ {
			if last[i] == first {
				// 相等,计数就可以
				num++
			} else {
				// 不相等,将上个记录的值和数写入buffer,更新新的值和数
				buffer.WriteString(strconv.Itoa(num))
				buffer.WriteString(string(first))
				first = last[i]
				num = 1
			}
		}
		// 将最后记录写入
		buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(string(first))
		return buffer.String()
	}
}
