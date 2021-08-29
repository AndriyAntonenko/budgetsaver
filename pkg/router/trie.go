package router

import (
	"strings"
)

type RouterTrie struct {
	Root     *RouterTrieNode
	hashSize int
}

func NewTrie(headValue string, hashSize int) *RouterTrie {
	return &RouterTrie{
		Root:     NewRouterTrieNode(headValue, hashSize),
		hashSize: hashSize,
	}
}

func (t *RouterTrie) AddNode(path string) {
	parts := strings.Split(path, "/")
	currentNode := t.Root

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		child := currentNode.Children.Lookup(part)
		if child == nil {
			child = NewRouterTrieNode(part, t.hashSize)
			currentNode.AddChild(child)
		}
		currentNode = child
	}

	currentNode.IsMethod = true
}

func (t *RouterTrie) Lookup(path string) bool {
	parts := strings.Split(path, "/")
	currentNode := t.Root

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		child := currentNode.Children.Lookup(part)
		if child == nil {
			return false
		}

		currentNode = child
	}

	return currentNode.IsMethod
}
