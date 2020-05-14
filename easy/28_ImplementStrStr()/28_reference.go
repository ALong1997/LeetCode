package leetcode

/*
解法: Rabin Karp - 滚动哈希 同时也是 Golang 源码中 strings.index 采用的算法

Michael O. Rabin和Richard M. Karp在1987年提出一个想法，即可以对模式串进行哈希运算并将其哈希值与文本中子串的哈希值进行比对。
这个思路有一个问题需要解决，如何在常数时间生成子串的哈希码？
滚动哈希：常数时间生成哈希码。生成一个长度为 L 数组的哈希码，需要 O(L) 时间。利用滑动窗口的特性，每次滑动都有一个元素进，一个出。

在这里我们将哈希简单当做128进制的字符串转换为10进制表示.
先将bbc计算出一个哈希, 就像16进制一样, 只不过这次是128"进制"才能表示asc字符的哈希 : b * 128^2 + b * 128^1 + c = 1605632 + 12544 + 99 = 1618275
然后用这个数值与父串比较, 如何比较呢? 我们来依次模拟这个过程
先从头计算bbb的哈希: b * 128^2 + b * 128^1 + b = 1618274 与 bbc是不等的, 那么就能得到bbb与bbc一定不等.
接下来往后移动一步: 还是bbb, 那是不是还得向上一步一样每个字符都去计算呢? 这里就有一个巧妙的地方了: 我们可以利用上一步已经计算好的值来简化本次计算
首先将第一个b减去: - b * 256^2, 然后向后移动一位: * 256 , 再加上最后一个b , 即: (1618275 - b * 128^2) * 128 + b = 1618274, 再与6447715(bbc)相比. 这种计算方式也叫"滚动哈希".
这样的比较方式与每个字符串比较的做法相比, 计算量固定为 -*+ 三个操作, 而不是子串长度的操作量. 所以当子串越长前者越有优势.
我们在移动一步: (1618274 - b * 128^2) * 128 + c = 1618275, 这时候就与1618274(bbc)相等了, 就能判定找到了bbc.

不过该算法到这还没有结束, 因为最后一步的判断可能会产生误差.
我们假设一个字符π的asc码为0, 字符Δ的asc码为12643, 那么字符串 bπΔ 的计算值就为 b * 128^2 + 0 + 12643 = 1618275 与bbc的值相等. 而明显bπΔ与bbc是不同的字符串, 所以为了避免这个错误, 所以当计算到两个字符串值相等的时候, 还需要再判断一次.
这种情况被称为哈希冲突, 不过这个情况很少见, 只要保证"进制"是足够大的素数. 在Golang中这个取值是16777619.


算法
计算子字符串 haystack.substring(0, L) 和 needle.substring(0, L) 的哈希值。
从起始位置开始遍历：从第一个字符遍历到第 N - L 个字符。
根据前一个哈希值计算滚动哈希。
如果子字符串哈希值与 needle 字符串哈希值相等，返回滑动窗口起始位置。
返回 -1，这时候 haystack 字符串中不存在 needle 字符串。

伪码
function RabinKarp(string s[1..n], string pattern[1..m])
  hpattern := hash(pattern[1..m]);
  for i from 1 to n-m+1
    hs := hash(s[i..i+m-1])
    if hs = hpattern
      if s[i..i+m-1] = pattern[1..m]
        return i
  return not found



结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/


// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func referenceStrStr(haystack string, needle string) int {
	//查找字符串的长度
	n := len(needle)
	switch {
	//长度 == 0， return 0
	case n == 0:
		return 0
	case n == 1:
		for k, _ := range haystack {
			if haystack[k] == needle[0] {
				return k
			}
		}
		return -1
	case n == len(haystack):
		if needle == haystack {
			return 0
		}
		return -1
	case n > len(haystack):
		return -1
	}
	// Rabin-Karp search

	hashsep, pow := hashStr(needle)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(haystack[i])
	}
	if h == hashsep && haystack[:n] == needle {
		return 0
	}
	for i := n; i < len(haystack); {
		h *= primeRK
		h += uint32(haystack[i])
		h -= pow * uint32(haystack[i-n])
		i++
		if h == hashsep && haystack[i-n:i] == needle {
			return i - n
		}
	}
	return -1
}

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
// hashStr也是也是1种rolling hash
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		// 核心算法是这一句(FNV Hash算法)
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	// pow = primeRK^(len(seq))
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
