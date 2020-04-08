package structure

// DoublyLinkedList is a golang implementation of the data structure
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert adds a new node at the beginning of the linked list
// O(1)
func (list *DoublyLinkedList) Insert(data int) *DoublyLinkedList {
	newNode := NewNode(data, list.head)
	list.head = newNode
	return list
}

// TailInsert adds a new node at the end of the doubly linked list
// O(1)
func (list *DoublyLinkedList) TailInsert(data int) {
	newNode := NewNode(data, nil)
	if list.head == nil {
		list.head = newNode
	}

	if list.tail != nil {
		list.tail.SetNextNode(newNode)
		list.tail = newNode
	}
}