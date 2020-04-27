package tree

type Iterator struct {
	tree Tree
}

func (iter *Iterator) Inorder(t Tree) []int {
	if t.Root() == nil {
		return nil
	}
	n := Node(t.Root())
	stack := NewStack()
	list := make([]int, 0)
	for ; ; {
		stack.Push(n)
		if n.Left() != nil {
			n = n.Left()
			continue
		}
		list = append(list, stack.Pop().Key())
		if n.Right() != nil {
			n = n.Right()
			continue
		}
	}
	return list
}
