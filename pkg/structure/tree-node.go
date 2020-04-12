package structure

// TreeNode is a golang implementation of a binary tree node
type TreeNode struct {
	data int
	LeftNode *TreeNode
	RightNode *TreeNode
	active bool
}

// Insert is a recursive method for inserting data into the BST
func (node *TreeNode) Insert(data int) {
	if data < node.data {
		if node.LeftNode == nil {
			node.LeftNode = NewTreeNode(data)
		} else {
			node.LeftNode.Insert(data)
		}
	} else {
		if node.RightNode == nil {
			node.RightNode = NewTreeNode(data)
		} else {
			node.RightNode.Insert(data)
		}
	}
}

// Find is a recursive method to find data stored in a BST
func (node *TreeNode) Find(data int) *TreeNode {
	if node.data == data && node.active {
		return node // Found the data
	}

	if data < node.data && node.LeftNode != nil {
		return node.LeftNode.Find(data) // Recursive call down to the left
	}
	
	if node.RightNode != nil {
		return node.RightNode.Find(data) // Recursive call to the right
	}

	return nil
}

// Min traverses the tree structure to the left recursively to find the smallest value in the tree
func (node *TreeNode) Min() int {
	if node.LeftNode == nil {
		return node.data
	}
	return node.LeftNode.Min()
}

// Max traverses the tree structure to the right recursively to find the largest value in the tree
func (node *TreeNode) Max() int {
	if node.RightNode == nil {
		return node.data
	}
	return node.RightNode.Max()
}

// TraverseInOrder performs an in-order tree traversal and prints the node data
func (node *TreeNode) TraverseInOrder() {
	if node.LeftNode != nil {
		node.LeftNode.TraverseInOrder()
	}

	println(node.data) // Do something to the current node here

	if node.RightNode != nil {
		node.RightNode.TraverseInOrder()
	}
}

// IsActive checks whether or not this node has been softly deleted
func (node *TreeNode) IsActive() bool {
	return node.active
}

// Deactivate the current node i.e. mark it as softly deleted
func (node *TreeNode) Deactivate() {
	node.active = false
}

// NewTreeNode factory method
func NewTreeNode(data int) *TreeNode {
	node := new(TreeNode)
	node.data = data
	node.LeftNode = nil
	node.RightNode = nil
	node.active = true

	return node
}