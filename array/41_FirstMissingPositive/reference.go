package LeetCode

/*
解法: 索引作为哈希表

数据预处理
	首先我们可以不考虑负数和零，因为不需要考虑。同样可以不考虑大于 n 的数字，
	因为首次缺失的正数一定小于或等于 n + 1 。
	缺失的正数为 n + 1 的情况会单独考虑。
	为了不考虑这些数，又要保证时间复杂度为 O(N) ，因此
	不能将这些元素弹出。我们可以将这些数用 1 替换。
	为了确保缺失的第一个正数不是 1，先要在这步操作前确定 1 是否存在。

如何实现就地算法
	现在我们有一个只包含正数的数组，范围为 1 到 n，
	现在的问题是在 O(N) 的时间和常数空间内找出首次缺失的正数。
	如果可以使用哈希表，且哈希表的映射是 正数 -> 是否存在 的话，这其实很简单。
	"脏工作环境" 的解决方法是将一个字符串 hash_str 分配 n 个 0，并且用类似于哈希表的方法，如果在数组中出现数字 i 则将字符串中 hash_str[i] 修改为 1 。
	我们不使用这种方法，但是借鉴这种 使用索引作为哈希键值 的想法。
	最终的想法是 使用索引作为哈希键 以及 元素的符号作为哈希值 来实现是否存在的检测。
	为了完成此操作，我们遍历一遍数组（该操作在数据预处理使得数组中只有正数的操作后），检查每个元素值 elem 以及将nums[elem] 元素的符号变为符号来表示数字 elem 出现在 nums 中。注意，当数字出现多次时需要保证符号只会变化 1 次。

算法
	检查 1 是否存在于数组中。如果没有，则已经完成，1 即为答案。将负数，零，和大于 n 的数替换为 1 。
	遍历数组。当读到数字 a 时，替换第 a 个元素的符号。
	注意重复元素：只能改变一次符号。由于没有下标 n ，使用下标 0 的元素保存是否存在数字 n。
	再次遍历数组(从 nums[1] 开始)。返回第一个正数元素的下标。
	如果 nums[0] > 0，则返回 n 。
	如果之前的步骤中没有发现 nums 中有正数元素，则返回n + 1。


结果: 执行用时 :0 ms 内存消耗 :2.2 MB
*/

func referenceFirstMissingPositive(nums []int) int {
	var length = len(nums)

	var flag bool
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			flag = true
		}
		if nums[i] <= 0 || nums[i] > length {
			nums[i] = 1
		}
	}
	if !flag {
		return 1
	}

	for i := 0; i < length; i++ {
		var a = nums[i]
		if a < 0 {
			a = -a
		}
		if a == length {
			nums[0] = -nums[0]
		} else if nums[a] > 0 {
			nums[a] = -nums[a]
		}
	}

	for i := 1; i < length; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return length
	} else {
		return length + 1
	}
}
