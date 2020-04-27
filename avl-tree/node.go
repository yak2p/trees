package avl_tree

import (
	"math"

	"github.com/yak2p/trees/tree"
)

type Node struct {
	key    int
	value  interface{}
	left   *Node
	right  *Node
	height int
}

func newNode(key int, value interface{}) *Node {
	return &Node{
		key:    key,
		value:  value,
		height: 1,
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

func (n *Node) getValue(key int) interface{} {
	if n.key == key {
		return n.value
	}
	if key < n.key && n.left != nil {
		return n.left.getValue(key)
	}
	if key > n.key && n.right != nil {
		return n.right.getValue(key)
	}
	return nil
}

func (n *Node) insert(key int, value interface{}) *Node {
	if key == n.key {
		// if the key already exist, just ignore
		return n
	}
	if key < n.key {
		if n.left == nil {
			n.insertLeft(key, value)
			return n
		}
		n.left = n.left.insert(key, value)
	} else {
		if n.right == nil {
			n.insertRight(key, value)
			return n
		}
		n.right = n.right.insert(key, value)
	}
	n.updateHeight()
	return n.rotateToBalance()
}

func (n *Node) rotateToBalance() *Node {
	switch n.balanceFactor() {
	case 2:
		if n.left.balanceFactor() >= 0 {
			return n.rightRotate()
		} else {
			n.left = n.left.leftRotate()
			return n.rightRotate()
		}
	case -2:
		if n.right.balanceFactor() <= 0 {
			return n.leftRotate()
		} else {
			n.right = n.right.rightRotate()
			return n.leftRotate()
		}
	default:
		return n
	}
}

func (n *Node) remove(key int) *Node {
	if n.key == key {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		leftBiggest := n.left
		for ; leftBiggest.right != nil; {
			leftBiggest = leftBiggest.right
		}
		n.key, n.value = leftBiggest.key, leftBiggest.value
		n.left = n.left.remove(leftBiggest.key)
	} else if key < n.key {
		if n.left == nil {
			return n
		}
		n.left = n.left.remove(key)
	} else {
		if n.right == nil {
			return n
		}
		n.right = n.right.remove(key)
	}
	n.updateHeight()
	return n.rotateToBalance()
}

func (n *Node) insertLeft(key int, value interface{}) {
	n.left = newNode(key, value)
	if n.height < 2 {
		n.height = 2
	}
}

func (n *Node) insertRight(key int, value interface{}) {
	n.right = newNode(key, value)
	if n.height < 2 {
		n.height = 2
	}
}

func (n *Node) balanceFactor() int {
	lh, rh := 0, 0
	if n.left != nil {
		lh = n.left.height
	}
	if n.right != nil {
		rh = n.right.height
	}
	return lh - rh
}

func (n *Node) isBalance() bool {
	return int(math.Abs(float64(n.balanceFactor()))) <= 1
}

// right rotate and return the new root node
func (n *Node) rightRotate() *Node {
	b := n.left
	br := b.right

	b.right = n
	n.left = br

	n.updateHeight()
	b.updateHeight()
	return b
}

// left rotate and return the new root node
func (n *Node) leftRotate() *Node {
	b := n.right
	bl := b.left

	b.left = n
	n.right = bl

	n.updateHeight()
	b.updateHeight()
	return b
}

func (n *Node) updateHeight() {
	lh, rh := 0, 0
	if n.left != nil {
		lh = n.left.height
	}
	if n.right != nil {
		rh = n.right.height
	}
	n.height = max(lh, rh) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
