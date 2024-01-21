package LeetCode

import (
	"github.com/ALong1997/LeetCode/common"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenSum := len(nums1) + len(nums2)
	mid := lenSum >> 1
	if lenSum&1 == 1 {
		return float64(binarySearchKth(nums1, nums2, mid+1))
	} else {
		return float64(binarySearchKth(nums1, nums2, mid)+binarySearchKth(nums1, nums2, mid+1)) / 2
	}
}

func binarySearchKth(nums1, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	if m > n {
		return binarySearchKth(nums2, nums1, k)
	}

	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return common.Min(nums1[0], nums2[0])
	}

	p1 := common.Min(k/2, m)
	p2 := k - p1
	if nums1[p1-1] < nums2[p2-1] {
		return binarySearchKth(nums1[p1:], nums2, k-p1)
	} else if nums1[p1-1] > nums2[p2-1] {
		return binarySearchKth(nums1, nums2[p2:], k-p2)
	} else {
		return nums1[p1-1]
	}
}
