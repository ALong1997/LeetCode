package leetcode

/*
解法: 双指针
我们注意到，只要 maxRight[i]>maxLeft[i] 积水高度将由 maxLeft 决定，类似地 maxLeft[i]>maxRight[i]
所以我们可以认为如果一端有更高的条形块（例如右端），积水的高度依赖于当前方向的高度（从左到右）。
当我们发现另一侧（右侧）的条形块高度不是最高的，我们则开始从相反的方向遍历（从右到左）。
我们必须在遍历时维护 maxLeft 和 maxRight ，但是我们现在可以使用两个指针交替进行，实现 1 次遍历即可完成。

算法
	初始化 left 指针为 0 并且 right 指针为 size-1
	While left<right, do:
		If height[left] < height[right]
			If height[left] > maxLeft, 更新 maxLeft
			Else 累加 maxLeft−height[left] 到 ans
			left = left + 1.
		Else
			If height[right] > maxRight, 更新 maxRight
			Else 累加 maxRight−height[right] 到 ans
			right = right - 1.


结果: 执行用时 :4 ms 内存消耗 :2.8 MB
*/

func reference2Trap(height []int) int {
	var ans int
	var length = len(height)
	var left, right, maxLeft, maxRight = 0, length - 1, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				ans += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				ans += maxRight - height[right]
			}
			right--
		}
	}
	return ans
}
