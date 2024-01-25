package LeetCode

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

/*
解法: 暴力解
对于数组中的每个元素，我们找出下雨后水能达到的最高位置，等于两边最大高度的较小值减去当前高度的值。

算法
	初始化 ans=0
	从左向右扫描数组：
		初始化 maxLeft=0 和 maxRight=0
		从当前元素向左扫描并更新：maxLeft=max(maxLeft,height[i])
		从当前元素向右扫描并更新：maxRight=max(maxRight,height[i])
		将 min(maxLeft,maxRight)−height[e] 累加到 ans


结果: 执行用时 :80 ms 内存消耗 :2.8 MB
*/

func trap(height []int) int {
	var ans int
	var maxLeft, maxRight int
	for e := range height {
		maxLeft, maxRight = 0, 0
		for i := e; i >= 0; i-- {
			if maxLeft < height[i] {
				maxLeft = height[i]
			}
		}
		for i := e; i < len(height); i++ {
			if maxRight < height[i] {
				maxRight = height[i]
			}
		}
		if maxLeft < maxRight {
			ans += maxLeft - height[e]
		} else {
			ans += maxRight - height[e]
		}
	}
	return ans
}
