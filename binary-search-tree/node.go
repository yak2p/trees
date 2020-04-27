package binary_search_tree

import (
	"github.com/yak2p/trees/tree"
)

type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
}

func newNode(key int, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (n *Node) Key() int {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Left() tree.Node {
	return n.left
}

func (n *Node) Right() tree.Node {
	return n.right
}

func (n *Node) insertRight(key int, value interface{}) {
	if n.right != nil {
		return
	}

	node := newNode(key, value)
	n.right = node
}

func (n *Node) insertLeft(key int, value interface{}) {
	if n.left != nil {
		return
	}

	node := newNode(key, value)
	n.left = node
}

func (n *Node) deleteLeft() {
	if n.left == nil {
		return
	}

	n.left = insteadChild(n.left)
}

func (n *Node) deleteRight() {
	if n.right == nil {
		return
	}

	n.right = insteadChild(n.right)
}

// return the node's child which can instead node if node deleted
func insteadChild(n *Node) *Node {
	if n.left == nil && n.right == nil {
		return nil
	}

	// choose the biggest node from left, replace value/key of n, return n
	if n.left != nil && n.right != nil {
		father := n
		biggest := father.left
		for ; ; {
			if biggest.right == nil {
				n.key = biggest.key
				n.value = biggest.value
				father.right = biggest.left
				return n
			}
			father = biggest
			biggest = biggest.right
		}
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	return nil
}
