package binary_tree

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{
			value: value,
			left:  nil,
			right: nil,
		}
	}
	insert(tree.root, value)
}

func insert(node *Node, value int) {
	if value > node.value {
		insertRight(node, value)
	} else if value < node.value {
		insertLeft(node, value)
	}
}

func insertLeft(node *Node, value int) {
	if node.left == nil {
		leftNode := &Node{
			value: value,
			left:  nil,
			right: nil,
		}
		node.left = leftNode
	} else {
		insert(node.left, value)
	}
}

func insertRight(node *Node, value int) {
	if node.right == nil {
		rightNode := &Node{
			value: value,
			left:  nil,
			right: nil,
		}
		node.right = rightNode
	} else {
		insert(node.right, value)
	}
}

func (tree *BinaryTree) Search(value int) bool {
	return internalSearch(tree.root, value)
}

func internalSearch(node *Node, value int) bool {
	if node == nil {
		return false
	} else if value == node.value {
		return true
	} else if value > node.value {
		return internalSearch(node.right, value)
	} else {
		return internalSearch(node.left, value)
	}
}

func (tree *BinaryTree) Remove(value int) {
	if tree.root != nil {
		if tree.root.value == value {
			tree.root = nil
		}
	}
}
