/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
*/
package leetcode

/*
解法: 滑动窗口

分别从0开始，从1开始...从len(words[0])-1开始,我们每次移动一个单词的长度，也就是 len(words[0]) 个字符，这样所有的移动被分成了三类。
情况一：当子串完全匹配，移动到下一个子串的时候。直接往后移动一个单词
情况二：当判断过程中，出现不符合的单词。直接移动到不满足的单词的下一个单词
情况三：判断过程中，出现的是符合的单词，但是次数超了。直接移动到子串中第一次出现超出次数的单词的下一个单词

结果: 执行用时 :48 ms 内存消耗 :7 MB
*/


func findSubstring(s string, words []string) []int {
	var ans []int
	var wordNum = len(words)
	if wordNum == 0 {
		return ans
	}

	var wordLen = len(words[0])
	// 记录 words 中单词出现次数
	var allWords = make(map[string]int, wordNum)
	for _, word := range words {
		allWords[word]++
	}

	//将所有移动分成 wordLen 类情况
	for j := 0; j < wordLen ; j++ {
		var hasWords = make(map[string]int, wordNum)
		var num = 0 //记录当前 hasWords 中有多少个单词
		//每次移动一个单词长度
		for i := j; i <= len(s) - wordNum * wordLen; i += wordLen {
			var hasRemoved = false
			LOOP:
				for num < wordNum {
					var word = s[i+num*wordLen : i+(num+1)*wordLen]
					var hasWord = false
					for k, v := range allWords {
						if k == word {
							hasWord = true
							hasWords[k]++
							//出现情况三，遇到了符合的单词，但是次数超了
							if hasWords[word] > v {
								hasRemoved = true
								var removeNum = 0
								//一直移除单词，直到次数符合了
								for hasWords[word] > v {
									var firstWord = s[i+removeNum*wordLen : i+(removeNum+1)*wordLen]
									hasWords[firstWord]--
									removeNum++
								}
								//加 1 是因为我们把当前单词加入到了 hasWords 中
								num = num - removeNum + 1
								//这里依旧是考虑到了最外层的 for 循环，看情况二的解释
								i = i + (removeNum - 1) * wordLen
								break LOOP
							}
							num++
						}
					}
					if !hasWord {
						//出现情况二，遇到了不匹配的单词，直接将 i 移动到该单词的后边（但其实这里
						//只是移动到了出现问题单词的地方，因为最外层有 for 循环， i 还会移动一个单词
						//然后刚好就移动到了单词后边）
						hasWords = make(map[string]int, wordNum)
						i = i + num * wordLen
						num = 0
						break
					}
				}
			if num == wordNum {
				ans = append(ans, i)
			}
			//出现情况一，子串完全匹配，我们将上一个子串的第一个单词从 hasWords 中移除
			if num > 0 && !hasRemoved {
				var firstWord = s[i : i+wordLen]
				hasWords[firstWord]--
				num = num - 1
			}
		}
	}
	return ans
}