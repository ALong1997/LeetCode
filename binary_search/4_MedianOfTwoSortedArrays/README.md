# [4.Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## Level
**Hard**


## Question

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

Example 2:

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

Constraints:

* nums1.length == m
* nums2.length == n
* 0 <= m <= 1000
* 0 <= n <= 1000
* 1 <= m + n <= 2000
* -106 <= nums1[i], nums2[i] <= 106

## Answer
### Solution

According to the definition of the median, when `m + n` is odd, the median is the `(m + n) / 2`-th element in the sorted arrays.
When `m + n` is even, the median is the average of the `(m + n) / 2`-th element and the `(m + n) / 2 + 1`-th element in the sorted arrays.
Therefore, this problem can be transformed into finding the k-th smallest number in two sorted arrays,
where k is `(m + n) / 2` or `(m + n) / 2 + 1`

The time complexity is O(log (m + n)), so binary search is adopted.

### Runtime
9 ms

### Memory Usage
4.67 MB


## Reference
### Solution

An array of length `m` can be divided into `m + 1` positions from `0` to `m`.
Only need to binary search for the cutting position in one array, the cutting position in the other array can be obtained accordingly.
In order to minimize the time complexity, we binary search the smaller of the two arrays.

First, generate a random `midA` by binary search. When does the cutting line meet the condition for the median?
That is, the numbers on the left side of the line are all smaller than the numbers on the right side, which means `nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA]`.
If these conditions are not met, the cutting line needs to be adjusted.
If `nums1[midA] < nums2[midB-1]`, it means that the numbers on the left side of the line defined by `midA` are too small, so the cutting line should be moved to the right.
If `nums1[midA-1] > nums2[midB]`, it means that the numbers on the left side of the line defined by `midA` are too large, so the cutting line should be moved to the left.

Assuming the cutting lines have been found, the indices of the two sides of array 1 with respect to the cutting line are `midA - 1` and `midA`.
The indices of the two sides of array 2 with respect to the cutting line are `midB - 1` and `midB`.
The final merged array, if the length is odd, the median will be `max(nums1[midA-1], nums2[midB-1])`.
If the length is even, the two numbers at the middle positions are `max(nums1[midA-1], nums2[midB-1])` and `min(nums1[midA], nums2[midB])`.
In that case, the median will be `(max(nums1[midA-1], nums2[midB-1]) + min(nums1[midA], nums2[midB])) / 2`.

### Runtime
4 ms

### Memory Usage
4.59 MB