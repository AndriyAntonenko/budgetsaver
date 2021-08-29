package router

import (
	"errors"
	"strings"
)

type RouterTrie struct {
	Root     *RouterTrieNode
	hashSize int
}

func NewRouterTrie(headValue string, hashSize int) *RouterTrie {
	return &RouterTrie{
		Root:     NewRouterTrieNode(headValue, hashSize),
		hashSize: hashSize,
	}
}

func (t *RouterTrie) AddNode(path string, handler Handler) {
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

	currentNode.Handler = &handler
}

func (t *RouterTrie) Lookup(path string) (Handler, error) {
	parts := strings.Split(path, "/")
	currentNode := t.Root

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		child := currentNode.Children.Lookup(part)
		if child == nil {
			return nil, errors.New("Handler not found")
		}

		currentNode = child
	}

	if currentNode.Handler == nil {
		return nil, errors.New("Handler not found")
	}

	return *currentNode.Handler, nil
}
