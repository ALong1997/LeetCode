package leetcode
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

import "math"

/*
   解法: 暴力解
   如果我们枚举「宽」，我们可以使用两重循环枚举矩形的左右边界以固定宽度 w，此时矩形的高度 h，就是所有包含在内的柱子的「最小高度」，对应的面积为 w * h。
   如果我们枚举「高」，我们可以使用一重循环枚举某一根柱子，将其固定为矩形的高度 h。随后我们从这跟柱子开始向两侧延伸，直到遇到高度小于 h 的柱子，就确定了矩形的左右边界。如果左右边界之间的宽度为 w，那么对应的面积为 w * h。

   结果: 执行用时 :612 ms 内存消耗 :4.2 MB
*/

// 枚举「宽」
func largestRectangleAreaWidth(heights []int) int {
	var maxArea int

	// 枚举左边界
	for left := 0; left < len(heights); left++ {
		var minHeight = math.MaxInt64
		// 枚举右边界
		for right := left; right < len(heights); right++ {
			// 确定高度
			if heights[right] == 0 {
				break
			}
			if heights[right] < minHeight {
				minHeight = heights[right]
			}
			// 计算面积
			if maxArea < (right-left+1)*minHeight {
				maxArea = (right - left + 1) * minHeight
			}
		}
	}

	return maxArea
}

// 枚举「高」
func largestRectangleAreaHeight(heights []int) int {
	var height, area, maxArea int

	for i := 0; i < len(heights); i++ {
		height = heights[i]
		area = height

		if height > 0 {
			// 向左延伸
			for j := i - 1; j >= 0; j-- {
				if heights[j] < height {
					break
				}
				area += height
			}

			// 向右延伸
			for j := i + 1; j < len(heights); j++ {
				if heights[j] < height {
					break
				}
				area += height
			}

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
