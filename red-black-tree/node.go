package red_black_tree

import "github.com/yak2p/trees/tree"

const (
	Red = iota
	Black
)

type Node struct {
	color               int
	key                 int
	value               interface{}
	left, right, parent *Node
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

func (n *Node) grandparent() *Node {
	if n.parent == nil {
		return nil
	}
	return n.parent.parent
}

func (n *Node) uncle() *Node {
	if n.parent == nil || n.parent.parent == nil {
		return nil
	}
	if n.parent == n.parent.parent.left {
		return n.parent.parent.right
	} else {
		return n.parent.parent.left
	}
}

func (n *Node) sibling() *Node {
	if n.parent == nil {
		return nil
	}
	if n.parent.left == n {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

func (n *Node) validate() (bool, int) {
	if n.color == Red {
		if n.parent == nil || n.parent.color == Red {
			return false, 0
		}
	}

	leftBlacks, rightBlacks := 0, 0
	valid := true
	if n.left != nil {
		valid, leftBlacks = n.left.validate()
		if !valid {
			return false, 0
		}
	}
	if n.right != nil {
		valid, rightBlacks = n.right.validate()
		if !valid {
			return false, 0
		}
	}
	if leftBlacks != rightBlacks {
		return false, 0
	}
	if n.color == Red {
		return true, leftBlacks
	} else {
		return true, leftBlacks + 1
	}
}

func (n *Node) biggestLeft() *Node {
	if n.left == nil {
		return nil
	}
	k := n.left
	for ; k.right != nil; {
		k = k.right
	}
	return k
}

func (n *Node) smallestRight() *Node {
	if n.right == nil {
		return nil
	}
	k := n.right
	for ; k.left != nil; {
		k = k.left
	}
	return k
}

func newNode(key int, value interface{}, parent *Node) *Node {
	return &Node{
		color:  Red,
		key:    key,
		value:  value,
		parent: parent,
	}
}
