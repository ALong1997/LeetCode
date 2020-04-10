/*
对于一个具有树特征的无向图，我们可选择任何一个节点作为根。
图因此可以成为树，在所有可能的树中，具有最小高度的树被称为最小高度树。
给出这样的一个图，写出一个函数找到所有的最小高度树并返回他们的根节点。
*/
package leetcode

/*
解法：DFS找出各节点作为根时的树高，再比较出最小值。
结果:超时。
*/

// root : 根节点
// haveSearch : 对应edges，该边是否查找过
// rootPath : 根节点到各顶点的路径长
func findMinHeightTrees(n int, edges [][]int) []int {
	haveSearch := make([]bool, n - 1)
	treeHeight := make([]int, n)
	for i := 0; i < n; i++ {
		// 初始化 haveSearch
		for j := 0; j < n - 1; j++ {
			haveSearch[j] = false
		}
		treeHeight[i] = dfsTreeHeight(i, haveSearch, edges)
	}
	// fmt.Printf("%v", treeHeight)
	return minHeight(treeHeight, n)
}

// 找出最小高度树并返回根节点数组
func minHeight(treeHeight []int, n int) (min []int) {
	if len(treeHeight) == 0 {
		return nil
	}
	// fmt.Printf("\n%v", treeHeight)
	minHeight := 65535
	for _, v := range treeHeight {
		if minHeight > v {
			minHeight = v
		}
	}
	for k, v := range treeHeight {
		if minHeight == v {
			min = append(min, k)
		}
	}
	return
}

// DFS查找以root为根节点的树高
func dfsTreeHeight(root int, haveSearch []bool, edges [][]int) int {
	height := 0
	for k, v := range edges {
		if haveSearch[k] {
			continue
		}
		if v[0] == root {
			haveSearch[k] = true
			temp := dfsTreeHeight(v[1], haveSearch, edges) + 1
			if height < temp {
				height = temp
			}
		} else if v[1] == root {
			haveSearch[k] = true
			temp := dfsTreeHeight(v[0], haveSearch, edges) + 1
			if height < temp {
				height = temp
			}
		}
	}
	return height
}