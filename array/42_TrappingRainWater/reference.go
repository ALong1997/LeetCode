package LeetCode

/*
解法: 动态编程
在暴力方法中，我们仅仅为了找到最大值每次都要向左和向右扫描一次。但是我们可以提前用数组存储这个值。因此，可以通过动态编程解决。

算法
	找到数组中从下标 e 到最左端最高的条形块高度 maxLeft。
	找到数组中从下标 e 到最右端最高的条形块高度 maxRight。
	扫描数组 height 并更新答案：
		累加 min(maxLeft[e],maxRight[e])−height[e] 到 ans 上


结果: 执行用时 :4 ms 内存消耗 :3.1 MB
*/

func referenceTrap(height []int) int {
	var ans int
	var length = len(height)
	if length == 0 {
		return ans
	}

	var maxLeft, maxRight = make([]int, length), make([]int, length)

	maxLeft[0] = height[0]
	for i := 1; i < length; i++ {
		if height[i] < maxLeft[i-1] {
			maxLeft[i] = maxLeft[i-1]
		} else {
			maxLeft[i] = height[i]
		}
	}

	maxRight[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		if height[i] < maxRight[i+1] {
			maxRight[i] = maxRight[i+1]
		} else {
			maxRight[i] = height[i]
		}
	}

	for i := 1; i < length-1; i++ {
		if maxLeft[i] < maxRight[i] {
			ans += maxLeft[i] - height[i]
		} else {
			ans += maxRight[i] - height[i]
		}
	}
	return ans
}
