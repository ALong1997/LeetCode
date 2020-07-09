/*
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:
单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。*/
package leetcode

/*
解法: 暴力解


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

func fullJustify(words []string, maxWidth int) []string {
	var ans []string

	for i := 0; i < len(words); {
		var s string
		var length, count = len(words[i]), 1
		for j := i+1; j < len(words); j++ {
			if length + len(words[j]) + count > maxWidth {
				break
			}
			length += len(words[j])
			count++
		}


		if i + count < len(words) {
			var blank = trim(maxWidth - length, count)
			for k := 0; k < count; k++ {
				s += words[i+k]
				if k == 0 || k != count-1 {
					for t := 0; t < blank[k]; t++ {
						s += " "
					}
				}
			}
		} else {
			// 文本最后一行
			for k := 0; k < count; k++ {
				s += words[i+k]
				maxWidth -= len(words[i+k])
				if maxWidth > 0 {
					s += " "
					maxWidth--
				}
			}
			for k := 0; k < maxWidth; k++ {
				s += " "
			}
		}

		ans = append(ans, s)
		i += count
	}

	return ans
}

// 计算空格排布
func trim(trimLength, count int) []int {
	if count == 1 {
		return []int{trimLength}
	}

	var blank = make([]int, count-1)
	var a, b, i = trimLength / (count-1), trimLength % (count-1), 0
	for ; i < b; i++ {
		blank[i] = a+1
	}
	for ; i < count-1; i++ {
		blank[i] = a
	}
	return blank
}