package router

type trieNodeType = int8

const (
	static trieNodeType = iota
	root
	pattern
)

type RouterTrieNode struct {
	path         string
	children     *RouterHashTable
	nodeType     trieNodeType
	patternChild bool
	handler      *Handler
}

func NewRouterTrieNode(path string, hashSize int, nodeType trieNodeType, patternChild bool) *RouterTrieNode {
	return &RouterTrieNode{
		path:     path,
		nodeType: nodeType,
		children: NewRouterHashTable(hashSize),
	}
}

func (tn *RouterTrieNode) addChild(node *RouterTrieNode) {
	tn.children.insert(node)
}
