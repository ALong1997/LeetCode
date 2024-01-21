package leetcode
/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。
但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
*/

/*
解法: 贪心算法


结果: 执行用时 :12 ms 内存消耗 :3.5 MB
*/

const (
	romanI  = 1    // [1, 4)
	romanIV = 4    // 4
	romanV  = 5    // [5, 9)
	romanIX = 9    // 9
	romanX  = 10   // [10, 40)
	romanXL = 40   // [40, 50)
	romanL  = 50   // [50, 90)
	romanXC = 90   // [90, 100)
	romanC  = 100  // [100, 400)
	romanCD = 400  // [400, 500)
	romanD  = 500  // [500, 900)
	romanCM = 900  // [900, 1000)
	romanM  = 1000 // [1000, maxInt)
)

func intToRoman(num int) string {
	var roman string
	for num > 0 {
		if num >= romanM {
			num -= romanM
			roman += "M"
			continue
		}
		if num >= romanCM {
			num -= romanCM
			roman += "CM"
			continue
		}
		if num >= romanD {
			num -= romanD
			roman += "D"
			continue
		}
		if num >= romanCD {
			num -= romanCD
			roman += "CD"
			continue
		}
		if num >= romanD {
			num -= romanD
			roman += "D"
			continue
		}
		if num >= romanC {
			num -= romanC
			roman += "C"
			continue
		}
		if num >= romanXC {
			num -= romanXC
			roman += "XC"
			continue
		}
		if num >= romanL {
			num -= romanL
			roman += "L"
			continue
		}
		if num >= romanXL {
			num -= romanXL
			roman += "XL"
			continue
		}
		if num >= romanX {
			num -= romanX
			roman += "X"
			continue
		}
		if num >= romanIX {
			num -= romanIX
			roman += "IX"
			continue
		}
		if num >= romanV {
			num -= romanV
			roman += "V"
			continue
		}
		if num >= romanIV {
			num -= romanIV
			roman += "IV"
			continue
		}
		if num >= romanI {
			num -= romanI
			roman += "I"
			continue
		}
	}
	return roman
}
