package structure

// Node is ia golang implementation of the Linked-List data structure
type Node struct {
	data int
	nextNode *Node
}

// SetData is a setter method for setting the data of a node in a linked-list
func (node *Node) SetData(data int) *Node {
	node.data = data
	return node
}

// GetData does what get data does
func (node *Node) GetData() int {
	return node.data
}

// SetNextNode is a setter method for attaching the next node in a linked-list
func (node *Node) SetNextNode(n *Node) *Node {
	node.nextNode = n
	return node
}

// NewNode is a factory function for creating/getting new nodes
func NewNode(data int, node *Node) *Node {
	newNode := new(Node).
		SetData(data).
		SetNextNode(node)

	return newNode
}
