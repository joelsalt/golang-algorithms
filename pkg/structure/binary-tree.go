package structure

// BinaryTree is a golang implementation of the data structure
type BinaryTree struct {
	root *TreeNode
}

// Insert adds a new data point to our BST
func (tree *BinaryTree) Insert(data int) *BinaryTree {
	if tree.root == nil {
		tree.root = NewTreeNode(data)
	} else {
		tree.root.Insert(data)
	}

	return tree
}

// Find is a search method that looks through our Binary Tree
func (tree *BinaryTree) Find(data int) *TreeNode {
	if tree.root != nil {
		return tree.root.Find(data)
	}

	return nil
}

// Min returns the smallest value in the tree
func (tree *BinaryTree) Min() int {
	return tree.root.Min()
}

// Max returns the latgest value in the tree
func (tree *BinaryTree) Max() int {
	return tree.root.Max()
}

// TraverseInOrder ...
func (tree *BinaryTree) TraverseInOrder() {
	if tree.root != nil {
		tree.root.TraverseInOrder()
	}
}

// SoftDelete is a logical deletion of a node with less complexity
func (tree *BinaryTree) SoftDelete(data int) {
	node := tree.Find(data)

	if node != nil {
		node.Deactivate()
	}
}

// Delete removes a node and chooses a successor
func (tree *BinaryTree) Delete(data int) {
	current := tree.root
	parent := tree.root
	isLeftChild := false

	for current != nil && current.data != data {
		parent = current

		if (data < current.data) {
			current = current.LeftNode
			isLeftChild = true
		} else {
			current = current.RightNode
			isLeftChild = false
		}
	}

	if current == nil { return }

	if current.LeftNode == nil && current.RightNode == nil { // Case 1, no child nodes
		if current == tree.root {
			tree.root = nil // delete root node
		} else {
			if isLeftChild {
				parent.LeftNode = nil
			} else {
				parent.RightNode = nil
			}
		}
	} else if current.RightNode == nil { // Case 2, only 1 child node
		if current == tree.root {
			tree.root = current.LeftNode
		} else if isLeftChild {
			parent.LeftNode = current.LeftNode
		} else {
			parent.RightNode = current.LeftNode
		}
	} else if current.LeftNode == nil { // Case 2, continued
		if current == tree.root {
			tree.root = current.RightNode
		} else if isLeftChild {
			parent.LeftNode = current.RightNode
		} else {
			parent.RightNode = current.RightNode
		}
	} else { // Case 3, node has 2 children
		// Find a successor node with a value slightly greater than the value to be deleted
		successor := getSuccessor(current)
		if current == tree.root { // if root is being deleted
			tree.root = successor
		} else if isLeftChild {
			parent.LeftNode = successor
		} else {
			parent.RightNode = successor
		}
		successor.LeftNode = current.LeftNode // the successor is now the parent of the deleted values children
	}
}

func getSuccessor(node *TreeNode) *TreeNode {
	parentOfSuccessor := node
	successor := node
	current := node.RightNode

	for current != nil { // Iterate over the tree
		parentOfSuccessor = successor
		successor = current
		current = current.LeftNode
	}

	if successor != node.RightNode {
		parentOfSuccessor.LeftNode = successor.RightNode
		successor.RightNode = node.RightNode
	}
	return successor
}