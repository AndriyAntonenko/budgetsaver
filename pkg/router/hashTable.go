package router

type RouterHashNode struct {
	value *RouterTrieNode
	next  *RouterHashNode
}

type RouterHashTable struct {
	table map[int]*RouterHashNode
	size  int
}

// Special index for patterns
const patterIndex = -1

func NewRouterHashTable(hashSize int) *RouterHashTable {
	if hashSize <= 0 {
		panic("hashSize should be positive")
	}

	return &RouterHashTable{
		size:  hashSize,
		table: map[int]*RouterHashNode{},
	}
}

func (hash *RouterHashTable) hashFunction(value string) int {
	hashKey := 0
	for charCode := range []byte(value) {
		hashKey = hashKey + charCode
	}

	return hashKey % hash.size
}

func (hash *RouterHashTable) insert(trieNode *RouterTrieNode) int {
	if trieNode.nodeType == pattern {
		element := RouterHashNode{value: trieNode, next: hash.table[patterIndex]}
		hash.table[patterIndex] = &element
		return patterIndex
	}
	index := hash.hashFunction(trieNode.path)
	element := RouterHashNode{value: trieNode, next: hash.table[index]}
	hash.table[index] = &element
	return index
}

func (hash *RouterHashTable) lookupAll(path string) *RouterTrieNode {
	if node := hash.lookupStatic(path); node != nil {
		return node
	}

	// Search for pattern
	if node := hash.lookupPattern(); node != nil {
		// TODO: Implement correct pattern search
		return node
	}

	return nil
}

func (hash *RouterHashTable) lookupPattern() *RouterTrieNode {
	if t := hash.table[patterIndex]; t != nil {
		// TODO: Implement correct pattern search
		return t.value
	}

	return nil
}

func (hash *RouterHashTable) lookupStatic(path string) *RouterTrieNode {
	index := hash.hashFunction(path)
	if t := hash.table[index]; t != nil {
		for t != nil {
			if t.value.path == path {
				return t.value
			}
			t = t.next
		}
	}

	return nil
}
