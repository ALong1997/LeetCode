package LeetCode

/*
解法: 创建HashMap，遍历数组的元素，如果target-nums[i]存在于map中，则返回i和map中对应的下标；否则把该元素存到map中。


结果: 执行用时 :4 ms 内存消耗 :4.6 MB
*/

// numsMap k是nums的value，v是nums的key组成的slice（因为nums内可能有重复值）
func referenceTwoSum(nums []int, target int) []int {
	var another int
	numsMap := make(map[int][]int)
	for k1, v1 := range nums {
		another = target - v1
		mapValue, ok := numsMap[another]
		if !ok {
			numsMap[v1] = []int{k1}
			continue
		}
		for _, v2 := range mapValue {
			return []int{v2, k1}
		}
	}
	return []int{0, 0}
}
