/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明：
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

提示：
	-10^9 <= nums1[i], nums2[i] <= 10^9
	nums1.length == m + n
	nums2.length == n
*/
package leetcode

/*
解法: 双指针 / 从后往前可以降低空间复杂度，不用再分配内存存储nums1


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	var p, q, r = m - 1, n - 1, m + n - 1
	for p >= 0 && q >= 0 && r >= 0 {
		if nums1[p] < nums2[q] {
			nums1[r] = nums2[q]
			q--
			r--
		} else {
			nums1[r] = nums1[p]
			p--
			r--
		}
	}

	if r < 0 || q < 0 {
		return
	} else if p <= 0 {
		for q >= 0 && r >= 0 {
			nums1[r] = nums2[q]
			q--
			r--
		}
	}
}
