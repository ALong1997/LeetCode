package LeetCode

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

/*
解法: 构建电话按键数字与字符对应的hash表,暴力解

结果: 执行用时 :0 ms 内存消耗 :2.1 MB
*/
func letterCombinations(digits string) []string {
	var ans []string
	// 电话按键数字与字符对应的hash表
	var phoneNum = map[byte][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return ans
	}

	ans = []string{""}
	for i := 0; i < len(digits); i++ {
		t := phoneNum[digits[i]-'0']
		var temp []string
		for j := 0; j < len(t); j++ {
			for z := 0; z < len(ans); z++ {
				temp = append(temp, ans[z]+t[j])
			}
		}
		ans = temp
	}
	return ans
}
