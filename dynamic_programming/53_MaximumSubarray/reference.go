package LeetCode

/*
解法: 分治
这个分治方法类似于「线段树求解 LCIS 问题」的 pushUp 操作。

我们定义一个操作 get(a, l, r) 表示查询 a 序列 [l,r] 区间内的最大子段和，那么最终我们要求的答案就是 get(nums, 0, nums.size() - 1)。
如何分治实现这个操作呢？对于一个区间 [l,r]，我们取 m = (l + r)/2，对区间 [l,m] 和 [m+1,r] 分治求解。
当递归逐层深入直到区间长度缩小为 1 的时候，递归「开始回升」。
这个时候我们考虑如何通过 [l,m] 区间的信息和 [m+1,r] 区间的信息合并成区间 [l,r] 的信息。最关键的两个问题是：
	我们要维护区间的哪些信息呢？
	我们如何合并这些信息呢？

对于一个区间 [l, r][l,r]，我们可以维护四个量：
	lSum 表示 [l,r] 内以 l 为左端点的最大子段和
	rSum 表示 [l,r] 内以 r 为右端点的最大子段和
	mSum 表示 [l,r] 内的最大子段和
	iSum 表示 [l,r] 的区间和

以下简称 [l,m] 为 [l,r] 的「左子区间」，[m+1,r] 为 [l,r] 的「右子区间」。
我们考虑如何维护这些量呢（如何通过左右子区间的信息合并得到 [l,r] 的信息）？对于长度为 1 的区间 [i,i]，四个量的值都和 ai 相等。对于长度大于 1 的区间：
	首先最好维护的是 iSum，区间 [l,r] 的 iSum 就等于「左子区间」的 iSum 加上「右子区间」的 iSum。
	对于 [l,r] 的 lSum，存在两种可能，它要么等于「左子区间」的 lSum，要么等于「左子区间」的 iSum 加上「右子区间」的 lSum，二者取大。
	对于 [l,r] 的 rSum，同理，它要么等于「右子区间」的 rSum，要么等于「右子区间」的 iSum 加上「左子区间」的 rSum，二者取大。
	当计算好上面的三个量之后，就很好计算 [l,r] 的 mSum 了。我们可以考虑 [l,r] 的 mSum 对应的区间是否跨越 m——它可能不跨越 m，也就是说 [l,r] 的 mSum 可能是「左子区间」的 mSum 和 「右子区间」的 mSum 中的一个；它也可能跨越 m，可能是「左子区间」的 rSum 和 「右子区间」的 lSum 求和。三者取大。


【题外话】
分治法相较于贪心法来说，时间复杂度相同，但是因为使用了递归，并且维护了四个信息的结构体，运行的时间略长，空间复杂度也不如方法一优秀，而且难以理解。那么这种方法存在的意义是什么呢？
对于这道题而言，确实是如此的。但是仔细观察「方法二」，它不仅可以解决区间 [0,n−1]，还可以用于解决任意的子区间 [l, r] 的问题。
如果我们把 [0,n−1] 分治下去出现的所有子区间的信息都用堆式存储的方式记忆化下来，即建成一颗真正的树之后，我们就可以在 O(logn) 的时间内求到任意区间内的答案。
我们甚至可以修改序列中的值，做一些简单的维护，之后仍然可以在 O(logn) 的时间内求到任意区间内的答案，对于大规模查询的情况下，这种方法的优势便体现了出来。这棵树就是上文提及的一种神奇的数据结构——线段树。

结果: 执行用时 :0 ms 内存消耗 :2 MB
*/

func referenceMaxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r status) status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) status {
	if l == r {
		return status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type status struct {
	lSum, rSum, mSum, iSum int
}
