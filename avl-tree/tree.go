package avl_tree

import (
	"github.com/yak2p/trees/tree"
)

type Tree struct {
	root *Node
}

func (t *Tree) Root() tree.Node {
	return t.root
}

func (t *Tree) GetValue(key int) interface{} {
	if t.root == nil {
		return nil
	}
	return t.root.getValue(key)
}

func (t *Tree) Insert(key int, value interface{}) {
	if t.root == nil {
		t.root = newNode(key, value)
	}
	t.root = t.root.insert(key, value)
}

func (t *Tree) Remove(key int) {
	if t.root == nil {
		return
	}
	t.root.remove(key)
}
