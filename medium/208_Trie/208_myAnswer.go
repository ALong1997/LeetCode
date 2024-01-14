package leetcode

/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
*/

/*
解法: 额外实现：
func (t *Trie) ChildrenCount(prefix string) uint32
func (t *Trie) Delete(prefix string) bool

结果: 执行用时 :32 ms 内存消耗 :30.39 MB
*/

type (
	Trie struct {
		root *node
	}

	node struct {
		children  [26]*node
		passCount uint32 // the number of words pass this node
		hasEnd    bool
	}
)

const letterStart = 'a'

func NewTrie() *Trie {
	return &Trie{root: &node{}}
}

// Insert the word into Trie
func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}

	current := t.root
	for _, letter := range word {
		current.passCount++
		childIndex := letter - letterStart
		if current.children[childIndex] == nil {
			current.children[childIndex] = &node{}
		}
		current = current.children[childIndex]
	}
	current.hasEnd = true
}

// Search return whether the word exists in Trie
func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	node := t.search(word)
	return node != nil && node.hasEnd
}

// Delete the word from Trie
func (t *Trie) Delete(word string) bool {
	if !t.Search(word) {
		return false
	}

	current := t.root
	for _, letter := range word {
		current.passCount--
		childIndex := letter - letterStart
		current = current.children[childIndex]
	}
	if current.passCount == 0 {
		current = nil
	} else {
		current.hasEnd = false
	}

	return true
}

// StartsWith return whether the prefix whether exists in Trie
func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	return t.search(prefix) != nil
}

// ChildrenCount return the number of words prefixed with prefix
func (t *Trie) ChildrenCount(prefix string) uint32 {
	node := t.search(prefix)
	if node == nil {
		return 0
	}

	return node.passCount
}

func (t *Trie) search(word string) *node {
	if t == nil || t.root == nil {
		return nil
	}

	if len(word) == 0 {
		return nil
	}

	current := t.root
	for _, letter := range word {
		childIndex := letter - letterStart
		if current.children[childIndex] == nil {
			return nil
		}
		current = current.children[childIndex]
	}

	return current
}
