package binary_search_tree

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) GetValue(key int) interface{} {
	if t.root == nil {
		return nil
	}

	return search(t.root, key)
}

func search(n *Node, key int) *Node {
	if n.Key() == key {
		return n
	}

	if key < n.Key() && n.left != nil {
		return search(n.left, key)
	}
	if key > n.Key() && n.right != nil {
		return search(n.right, key)
	}
	return nil
}

func (t *BinarySearchTree) Insert(key int, value interface{}) {
	if t.root == nil {
		t.root = newNode(key, value)
		return
	}

	current := t.root
	for ; ; {
		if key == current.Key() {
			return
		}
		if key < current.Key() {
			if current.left == nil {
				current.insertLeft(key, value)
				return
			}
			current = current.left
			continue
		}
		if key > current.Key() {
			if current.right == nil {
				current.insertRight(key, value)
				return
			}
			current = current.right
			continue
		}
	}
}

func (t *BinarySearchTree) InsertRecursive(key int, value interface{}) {
	if t.root == nil {
		t.root = newNode(key, value)
		return
	}

	insert(t.root, key, value)
}

func insert(n *Node, key int, value interface{}) {
	if key == n.Key() {
		return
	}
	if key < n.Key() {
		if n.left != nil {
			insert(n.left, key, value)
		} else {
			n.insertLeft(key, value)
		}
		return
	} else {
		if n.right != nil {
			insert(n.right, key, value)
		} else {
			n.insertRight(key, value)
		}
	}
}

func (t *BinarySearchTree) Remove(key int) {
	if t.root == nil {
		return
	}

	t.root = remove(t.root, key)
}

func remove(n *Node, key int) *Node {
	if key == n.Key() {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		leftBiggest := n.left
		for {
			if leftBiggest.right != nil {
				leftBiggest = leftBiggest.right
			}
			break
		}
		n.key, n.value = leftBiggest.key, leftBiggest.value
		n.left = remove(n.left, leftBiggest.key)
		return n
	}
	if key < n.Key() && n.left != nil {
		n.left = remove(n.left, key)
	}
	if key > n.Key() && n.right != nil {
		n.right = remove(n.right, key)
	}
	return n
}
