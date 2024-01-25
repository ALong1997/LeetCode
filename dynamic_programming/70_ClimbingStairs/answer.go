package LeetCode

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/

/*
解法: 动态规划
我们用 f(x) 表示爬到第 x 级台阶的方案数，考虑最后一步可能跨了一级台阶，也可能跨了两级台阶，所以我们可以列出如下式子： f(x) = f(x-1) + f(x-2)
它意味着爬到第 x 级台阶的方案数是爬到第 x - 1 级台阶的方案数和爬到第 x−2 级台阶的方案数的和。

可以用「滚动数组思想」把空间复杂度优化成 O(1)O(1)。下面的代码中给出的就是这种实现。


结果: 执行用时 :0 ms 内存消耗 :1.9 MB
*/

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
