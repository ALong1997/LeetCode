/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
*/
package leetcode

/*
解法: 对于两个有序数组来说，我们只要找出第 (m+n+1)/2 大的数和第 (m+n+2)/2 大的数，然后求平均数即可。
要求算法的时间复杂度为 O(log(m + n))故采用二分查找


结果: 执行用时 :12 ms 内存消耗 :5.6 MB
*/
const maxInt = int(^uint(0) >> 1)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k1 := (m + n + 1) / 2
	k2 := (m + n + 2) / 2
	return (float64(getKth(nums1, nums2, m, n, 0, 0, k1) + getKth(nums1, nums2, m, n, 0, 0, k2))) / 2.0
}

// 二分查找有序数组第 k 大的数
func getKth(nums1, nums2 []int, m, n, left1, left2, k int) int {
	// 当某个数组查找的起始位置大于等于该数组长度时，说明这个数组中的所有数已经被淘汰，则只需要在另一个数组找查找即可。
	if left1 > m - 1 {
		return nums2[left2 + k - 1]
	}
	if left2 > n - 1 {
		return nums1[left1 + k - 1]
	}
	// 如果 k = 1，即需要查找第一个数。
	if k == 1 {
		if nums1[left1] < nums2[left2] {
			return nums1[left1]
		}
		return nums2[left2]
	}

	// 分别在两个数组中查找第 k/2 个元素，若存在（即数组没有越界），标记为找到的值；若不存在，标记为整数最大值
	mid1, mid2 := maxInt, maxInt
	if left1 + k/2 - 1 < m {
		mid1 = nums1[left1 + k/2 - 1]
	}
	if left2 + k/2 - 1 < n {
		mid2 = nums2[left2 + k/2 - 1]
	}
	// 确定最终的第k/2个元素，然后递归查找
	if mid1 < mid2 {
		return getKth(nums1, nums2, m, n, left1 + k/2, left2, k - k/2)
	}
	return getKth(nums1, nums2, m, n, left1, left2 + k/2, k - k/2)
}