package LeetCode

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
*/

/*
解法: 动态规划，双指针
所求即找出最大的 area = (right-left)*min(a[left], a[right]) (left < right)
显然地，当 left < i < right 时，如若 a[i] < min(a[left], a[right]),则由 left 与 i 或 right 与 i 之间容纳的水必少于 left 与 right 之间能容纳的水量

在初始时，左右指针分别指向数组的左右两端
此时我们需要移动一个指针。移动哪一个呢？直觉告诉我们，应该移动对应数字较小的那个指针（即此时的左指针）。
这是因为，由于容纳的水量是由两个指针指向的数字中较小值 * 指针之间的距离决定的。
如果我们移动数字较大的那个指针，那么前者「两个指针指向的数字中较小值」不会增加，后者「指针之间的距离」会减小，那么这个乘积会减小。
因此，我们移动数字较大的那个指针是不合理的。因此，我们移动 数字较小的那个指针。

结果: 执行用时 :16 ms 内存消耗 :5.8 MB
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	areaMax := 0
	for left < right {
		min := height[left]
		if min > height[right] {
			min = height[right]
		}
		tempArea := (right - left) * min
		if areaMax < tempArea {
			areaMax = tempArea
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return areaMax
}
