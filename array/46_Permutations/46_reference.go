package leetcode

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
*/

/*
解法: 回溯
这个问题可以看作有 n 个排列成一行的空格，我们需要从左往右依此填入题目给定的 n 个数，每个数只能使用一次。
那么很直接的可以想到一种穷举的算法，即从左往右每一个位置都依此尝试填入一个数，看能不能填完这 n 个空格，在程序中我们可以用「回溯法」来模拟这个过程。

我们定义递归函数 backtrack(first, output) 表示从左往右填到第 first 个位置，当前排列为 output。 那么整个递归函数分为两个情况：
	如果 first==n，说明我们已经填完了 n 个位置（注意下标从 0 开始），找到了一个可行的解，我们将 output 放入答案数组中，递归结束。
	如果 first<n，我们要考虑这第 first 个位置我们要填哪个数。根据题目要求我们肯定不能填已经填过的数。
		因此很容易想到的一个处理手段是我们定义一个标记数组 vis[] 来标记已经填过的数，那么在填第 first 个数的时候我们遍历题目给定的 n 个数，如果这个数没有被标记过，我们就尝试填入，并将其标记，继续尝试填下一个位置，即调用函数 backtrack(first + 1, output)。
		搜索回溯的时候要撤销这一个位置填的数以及标记，并继续尝试其他没被标记过的数。

使用标记数组来处理填过的数是一个很直观的思路，但是可不可以去掉这个标记数组呢？毕竟标记数组也增加了我们算法的空间复杂度。
我们可以将题目给定的 n 个数的数组 nums[] 划分成左右两个部分，左边的表示已经填过的数，右边表示待填的数，我们在递归搜索的时候只要动态维护这个数组即可。
具体来说，假设我们已经填到第 first 个位置，那么 nums[] 数组中 [0,first−1] 是已填过的数的集合，[first,n−1] 是待填的数的集合。
我们肯定是尝试用 [first,n−1] 里的数去填第 first 个数，假设待填的数的下标为 i ，那么填完以后我们将第 i 个数和第 first 个数交换，即能使得在填第 first+1个数的时候 nums[] 数组的 [0,first] 部分为已填过的数，[first+1,n−1] 为待填的数，回溯的时候交换回来即能完成撤销操作。

当然善于思考的读者肯定已经发现这样生成的全排列并不是按字典序存储在答案数组中的，如果题目要求按字典序输出，那么请还是用标记数组或者其他方法。

结果: 执行用时 :0 ms 内存消耗 :2.6 MB
*/

func permute(nums []int) [][]int {
	var ans [][]int
	var length = len(nums)
	var backtrack func(int)
	backtrack = func(first int) {
		if first == length {
			var temp = make([]int, length)
			copy(temp, nums)
			ans = append(ans, temp)
		}
		for i := first; i < length; i++ {
			// 动态维护数组
			nums[first], nums[i] = nums[i], nums[first]
			// 继续递归填下一个数
			backtrack(first + 1)
			// 撤销操作
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return ans
}
