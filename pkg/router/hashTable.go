package router

type RouterHashNode struct {
	Value *RouterTrieNode
	Next  *RouterHashNode
}

type RouterHashTable struct {
	Table map[int]*RouterHashNode
	Size  int
}

func NewRouterHashTable(hashSize int) *RouterHashTable {
	return &RouterHashTable{
		Size:  hashSize,
		Table: map[int]*RouterHashNode{},
	}
}

func (hash *RouterHashTable) hashFunction(value string) int {
	hashKey := 0
	for charCode := range []byte(value) {
		hashKey = hashKey + charCode
	}

	return hashKey & hash.Size
}

func (hash *RouterHashTable) Insert(trieNode *RouterTrieNode) int {
	index := hash.hashFunction(trieNode.Value)
	element := RouterHashNode{Value: trieNode, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func (hash *RouterHashTable) Lookup(value string) *RouterTrieNode {
	index := hash.hashFunction(value)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value.Value == value {
				return t.Value
			}
			t = t.Next
		}
	}

	return nil
}
