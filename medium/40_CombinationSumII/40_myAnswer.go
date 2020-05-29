/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。

说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
*/
package leetcode

import (
	"sort"
)

/*
解法: 回溯 + 剪枝 基于T39
本题与T39区别在于 candidates 中可能有重复值，且 candidates 中的每个数字在每个组合中只能使用一次。

编码的不同在于下一层递归的起始索引不一样。
	第 39 题：还从候选数组的当前索引值开始。
	第 40 题：从候选数组的当前索引值的下一位开始。
为了使得解集不包含重复的组合。
我们想一想，由于已对数组升序排序，重复的元素一定不是排好序以后的第 1 个元素和相同元素的第 1 个元素。
候选数组有序，对于在递归树中发现重复分支，进而“剪枝”也是有效的。

结果: 执行用时 :4 ms 内存消耗 :2.4 MB
*/


func combinationSum2(candidates []int, target int) [][]int {
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
		if i != begin && candidates[i] == candidates[i-1] { // *同层节点 数值相同，剪枝
			continue
		}
		value = target - candidates[i]
		// “剪枝”操作，不必递归到下一层，并且后面的分支也不必执行
		if value < 0 {
			break
		}
		path = append(path, candidates[i])
		// 因为下一层不能比上一层还小，起始索引还从 i+1 开始
		dfs(candidates, path, i+1, length, value, ans)
		// 这一步可能造成指针污染
		path = path[:len(path)-1]
	}
}
