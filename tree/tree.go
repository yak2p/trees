package tree

type Node interface {
	Key() int
	Value() interface{}
	Left() Node
	Right() Node
}

type BinaryTree interface {
	// get the value of the key, if the node is not exist, return nil
	GetValue(key int) interface{}

	// insert a node with key if not exist
	Insert(key int, value interface{})

	// delete node from tree if exist
	Remove(key int)

	//return the root node
	Root() Node
}
