package red_black_tree

import "github.com/yak2p/trees/tree"

type Tree struct {
	root *Node
}

func (t *Tree) GetValue(key int) interface{} {
	return t.search(key)
}

func (t *Tree) Insert(key int, value interface{}) {
	t.insert(key, value)
}

func (t *Tree) Remove(key int) {
	if t.root == nil {
		return
	}
	n := t.root
	for {
		if n.key == key {
			t.remove(n)
			return
		}
		if key < n.key && n.left != nil {
			n = n.left
			continue
		}
		if key > n.key && n.right != nil {
			n = n.right
			continue
		}
		break
	}
}

func (t *Tree) Root() tree.Node {
	return t.root
}

func (t *Tree) search(key int) interface{} {
	if t.root == nil {
		return nil
	}
	n := t.root
	for {
		if n.key == key {
			return n.value
		}
		if key < n.key && n.left != nil {
			n = n.left
			continue
		}
		if key > n.key && n.right != nil {
			n = n.right
			continue
		}
		return nil
	}
}

func (t *Tree) insert(key int, value interface{}) {
	if t.root == nil {
		t.root = &Node{
			color: Black,
			key:   key,
			value: value,
		}
		return
	}
	n := t.root
	for {
		// insert a already existent key means update
		if key == n.key {
			n.value = value
			return
		}
		if key < n.key {
			if n.left == nil {
				n.left = newNode(key, value, n)
				t.insertRebalance(n.left)
				return
			} else {
				n = n.left
			}
		} else {
			if n.right == nil {
				n.right = newNode(key, value, n)
				t.insertRebalance(n.right)
				return
			} else {
				n = n.right
			}
		}
	}
}

func (t *Tree) insertRebalance(n *Node) {
	// Case 1
	if n.parent == nil {
		n.color = Black
		return
	}
	// Case 2
	if n.parent.color == Black {
		return
	}
	// Case 3
	if n.uncle() != nil && n.uncle().color == Red {
		n.parent.color = Black
		n.uncle().color = Black
		n.grandparent().color = Red
		t.insertRebalance(n.grandparent())
		return
	}
	// Case 4 if n's parent is Red, n must have grandparent
	if n.parent == n.grandparent().left {
		if n == n.parent.right {
			n.grandparent().color = Red
			n.color = Black
			t.leftRotate(n)
			t.rightRotate(n)
		} else {
			n.grandparent().color = Red
			n.parent.color = Black
			t.rightRotate(n.parent)
		}
	} else {
		if n == n.parent.left {
			n.grandparent().color = Red
			n.color = Black
			t.rightRotate(n)
			t.leftRotate(n)
		} else {
			n.grandparent().color = Red
			n.parent.color = Black
			t.leftRotate(n.parent)
		}
	}
}

func (t *Tree) leftRotate(n *Node) {
	parent := n.parent
	if n.grandparent() != nil {
		if n.parent == n.grandparent().left {
			n.grandparent().left = n
		} else {
			n.grandparent().right = n
		}
	} else {
		// N's grandparent is nil means n's parent is root
		t.root = n
	}
	n.parent = n.grandparent()

	parent.right = n.left
	if parent.right != nil {
		parent.right.parent = parent
	}

	n.left = parent
	parent.parent = n
}

func (t *Tree) rightRotate(n *Node) {
	parent := n.parent
	if n.grandparent() != nil {
		if n.parent == n.grandparent().left {
			n.grandparent().left = n
		} else {
			n.grandparent().right = n
		}
	} else {
		t.root = n
	}
	n.parent = n.grandparent()

	parent.left = n.right
	if parent.left != nil {
		parent.left.parent = parent
	}

	n.right = parent
	parent.parent = n
}

func (t *Tree) remove(n *Node) {
	if n.left != nil && n.right != nil {
		instead := n.biggestLeft()
		n.key = instead.key
		n.value = instead.value
		t.remove(instead)
		return
	}
	child := n.left
	if child == nil {
		child = n.right
	}
	if n.color == Red || (child != nil && child.color == Red) {
		if child != nil {
			child.color = Black
			child.parent = n.parent
		}
		if n.parent != nil {
			if n.parent.left == n {
				n.parent.left = child
			} else {
				n.parent.right = child
			}
		} else {
			t.root = child
		}
		return
	}
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	} else {
		t.root = nil
	}
	if n.left == nil && n.right == nil {
		t.removeRebalance(n.parent, nil)
	}
}

func isBlack(n *Node) bool {
	return n == nil || n.color == Black
}

func (t *Tree) removeRebalance(p *Node, n *Node) {
	// Case 1
	if p == nil {
		return
	}
	// N is black (even when n is nil), so n must has sibling if n has parent
	s := p.left
	if p.left == n {
		s = p.right
	}
	// Case 2
	if p.color == Red {
		// Case 2.2.1
		if isBlack(s.left) && isBlack(s.right) {
			p.color = Black
			s.color = Red
			return
		}
		// Case 2.2.2
		if p.left == n {
			// Case 2.2.2.1
			if isBlack(s.right) {
				s.left.color = Black
				s.color = Red
				t.rightRotate(s.left)
				t.removeRebalance(p, n)
				return
			}
			// Case 2.2.2.2
			s.color = Red
			p.color = Black
			s.right.color = Black
			t.leftRotate(s)
			return
		} else {
			// Case 2.2.3
			// Case 2.2.3.1
			if isBlack(s.left) {
				s.right.color = Black
				s.color = Red
				t.leftRotate(s.right)
				t.removeRebalance(p, n)
				return
			}
			// Case 2.2.3.2
			s.color = Red
			p.color = Black
			s.left.color = Black
			t.rightRotate(s)
			return
		}
	}
	// Case 3 p.color == Black
	// Case 3.1
	if s.color == Red {
		s.color = Black
		p.color = Red
		if n == p.left {
			t.leftRotate(s)
		} else {
			t.rightRotate(s)
		}
		t.removeRebalance(p, n)
		return
	}
	// Case 3.2 s.color == Black
	// Case 3.2.1
	if isBlack(s.left) && isBlack(s.right) {
		s.color = Red
		t.removeRebalance(p.parent, p)
		return
	}
	// Case 3.2.2
	if p.left == n {
		// Case 3.2.2.1
		if isBlack(s.right) {
			s.left.color = Black
			s.color = Red
			t.rightRotate(s.left)
			t.removeRebalance(p, n)
			return
		}
		// Case 3.2.2.2
		s.right.color = Black
		t.leftRotate(s)
		return
	} else {
		// Case 3.2.3
		// Case 3.2.3.1
		if isBlack(s.left) {
			s.right.color = Black
			s.color = Red
			t.leftRotate(s.right)
			t.removeRebalance(p, n)
			return
		}
		// Case 3.2.3.2
		s.left.color = Black
		t.rightRotate(s)
		return
	}
}

func (t *Tree) validate() bool {
	if t.root == nil {
		return true
	} else {
		valid, _ := t.root.validate()
		return valid
	}
}
