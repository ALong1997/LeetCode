/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。

说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
*/
package leetcode

import (
	"sort"
)

/*
解法: 回溯 + 剪枝
尝试找到组合之和为该数的所有组合，怎么找呢？逐个减掉候选数组中的元素即可；
以 candidates = [2,3,6,7], target = 7 为例, target = 7 为根结点，每一个分支做减法；
减到 0 或者负数的时候，到了叶子结点；
减到 0 的时候结算，这里 “结算” 的意思是添加到结果集；
从根结点到叶子结点（必须为 0）的路径，就是题目要我们找的一个组合。
尝试画出这棵树。
发现可能有重复结果:重复的原因是在较深层的结点值考虑了之前考虑过的元素，因此我们需要设置“下一轮搜索的起点”即可（这里可能没有说清楚，已经尽力了）。

去重复
	在搜索的时候，需要设置搜索起点的下标 begin ，由于一个数可以使用多次，下一层的结点从这个搜索起点开始搜索；
	在搜索起点 begin 之前的数因为以前的分支搜索过了，所以一定会产生重复。

剪枝提速
	如果一个数位搜索起点都不能搜索到结果，那么比它还大的数肯定搜索不到结果，基于这个想法，我们可以对输入数组进行排序，以减少搜索的分支；
	排序是为了提高搜索速度，非必要；
	搜索问题一般复杂度较高，能剪枝就尽量需要剪枝。把候选数组排个序，遇到一个较大的数，如果以这个数为起点都搜索不到结果，后面的数就更搜索不到结果了。


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/


func combinationSum(candidates []int, target int) [][]int {
	var length = len(candidates)
	var ans [][]int
	if length == 0 {
		return ans
	}

	// 剪枝是为了提高搜索速度，非必要；
	sort.Ints(candidates)
	// 在遍历的过程中记录路径，它是一个栈
	var path []int
	dfs(candidates, path, 0, length, target, &ans)
	return ans
}

func dfs(candidates, path []int, begin, length, target int, ans *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		// 深拷贝，保证已加入结果集不受影响
		// path作为参数，是一直在变化的，如果不拷贝的话，ans中的所有元素底层都是共用同一个数组(path)了
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	var value int //差值
	for i := begin; i < length; i++ {
		value = target - candidates[i]
		// “剪枝”操作，不必递归到下一层，并且后面的分支也不必执行
		if value < 0 {
			break
		}
		path = append(path, candidates[i])
		// 因为下一层不能比上一层还小，起始索引还从 i 开始
		dfs(candidates, path, i, length, value, ans)
		// 这一步可能造成指针污染
		path = path[1:]
	}
}
