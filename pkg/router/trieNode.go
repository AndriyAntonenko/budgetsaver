package router

type RouterTrieNode struct {
	Value    string
	Children *RouterHashTable
	Handler  *Handler
}

func NewRouterTrieNode(value string, hashSize int) *RouterTrieNode {
	return &RouterTrieNode{
		Value:    value,
		Children: NewRouterHashTable(hashSize),
	}
}

func (tn *RouterTrieNode) AddChild(node *RouterTrieNode) {
	tn.Children.Insert(node)
}
