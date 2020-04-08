package structure

// DoublyLinkedNode is a golang implimentation of a doubly linked list
type DoublyLinkedNode struct {
	data int
	PreviousNode *DoublyLinkedNode
	NextNode *DoublyLinkedNode
}

// SetData is a setter method for setting the data of a node in a linked-list
func (node *DoublyLinkedNode) SetData(data int) *DoublyLinkedNode {
	node.data = data
	return node
}

// GetData does what get data does
func (node *DoublyLinkedNode) GetData() int {
	return node.data
}

// SetNextNode is a setter method for attaching the next node in a linked-list
func (node *DoublyLinkedNode) SetNextNode(n *DoublyLinkedNode) *DoublyLinkedNode {
	node.NextNode = n
	return node
}

// SetPreviousNode is a setter method for attaching the next node in a linked-list
func (node *DoublyLinkedNode) SetPreviousNode(n *DoublyLinkedNode) *DoublyLinkedNode {
	node.PreviousNode = n
	return node
}

// NewDoublyLinkedNode is a factory function for creating/getting new nodes
func NewDoublyLinkedNode(
	data int,
	nextNode *DoublyLinkedNode,
	previousNode *DoublyLinkedNode,
) *DoublyLinkedNode {
	newNode := new(DoublyLinkedNode).
		SetData(data).
		SetNextNode(nextNode).
		SetPreviousNode(previousNode)

	return newNode
}