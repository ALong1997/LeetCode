package LeetCode

/*
解法: 单调栈（基于 枚举「高」 的优化）



结果: 执行用时 :12 ms 内存消耗 :6 MB
*/

func referenceLargestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	var monoStack []int
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
