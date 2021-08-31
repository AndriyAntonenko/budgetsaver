package router

import (
	"fmt"
	"strings"
)

type RouterTrie struct {
	root     *RouterTrieNode
	hashSize int
}

func NewRouterTrie(headValue string, hashSize int) *RouterTrie {
	return &RouterTrie{
		root:     NewRouterTrieNode(headValue, hashSize, root, false),
		hashSize: hashSize,
	}
}

func (t *RouterTrie) AddNode(path string, handler Handler) {
	parts := strings.Split(path, "/")
	currentNode := t.root

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		// TODO: Avoid pattern match
		var child *RouterTrieNode

		_, isParam := getParam(part)
		if isParam {
			child = currentNode.children.lookupPattern()
		} else {
			child = currentNode.children.lookupStatic(part)
		}

		if child == nil {
			patternChild := currentNode.nodeType == pattern
			if isParam {
				child = NewRouterTrieNode(part, t.hashSize, pattern, patternChild)
			} else {
				child = NewRouterTrieNode(part, t.hashSize, static, patternChild)
			}
			fmt.Printf("Is %s pattern child? %v. Parent value: %s\n", child.path, patternChild, currentNode.path)

			currentNode.addChild(child)
		}

		currentNode = child
	}

	currentNode.handler = &handler
}

func (t *RouterTrie) Lookup(path string) Handler {
	parts := strings.Split(path, "/")
	currentNode := t.root

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		child := currentNode.children.lookupAll(part)
		if child == nil {
			return nil
		}

		currentNode = child
	}

	if currentNode.handler == nil {
		return nil
	}

	// fmt.Printf("Is %s pattern child? %v\n", currentNode.path, currentNode.patternChild)
	return *currentNode.handler
}

func getParam(part string) (paramName string, isParam bool) {
	firstChar := part[0]
	paramName = ""
	isParam = false

	if string(firstChar) == ":" {
		paramName = string(part[1:])
		isParam = true
	}

	return
}
