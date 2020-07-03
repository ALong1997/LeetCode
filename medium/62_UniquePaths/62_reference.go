package leetcode

/*
解法: 动态规划
我们令 dp[i][j] 是到达 i, j 最多路径
动态方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]

注意，对于第一行 dp[0][j]，或者第一列 dp[i][0]，由于都是在边界，所以只能为 1

优化每次仅用到 dp[i-1][j] 和 dp[i][j-1] 所以可将dp优化为一个一维数组中:cur[j] += cur[j-1]
     a) cur 长度取 n 防止溢出
     b) cur[i]每行遍历只更新一次,也就是说cur[j]原有的值代表的是上一行dp[i-1],[j]的值,而cur[j-1] 是在本行遍历刚更新过的值因此是dp[i][j-1]的值!



结果: 执行用时 :0 ms 内存消耗 :1.9 MB
*/


func referenceUniquePaths(m int, n int) int {
	cur := make([]int,n)
	for i:=0;i<n;i++{
		cur[i] = 1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}