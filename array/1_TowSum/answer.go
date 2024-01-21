package LeetCode

/*
   解法: 首尾递进查找


   Result: 执行用时 :0 ms 内存消耗 :4.02 MB
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
