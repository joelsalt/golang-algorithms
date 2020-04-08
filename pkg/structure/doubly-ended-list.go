package structure

// DoublyEndedList is a golang implementation of the data structure
type DoublyEndedList struct {
	head *Node
	tail *Node
}

// Insert adds a new node at the beginning of the linked list
// O(1)
func (list *DoublyEndedList) Insert(data int) *DoublyEndedList {
	newNode := NewNode(data, list.head)
	list.head = newNode
	return list
}

// TailInsert adds a new node at the end of the doubly linked list
// O(1)
func (list *DoublyEndedList) TailInsert(data int) {
	newNode := NewNode(data, nil)
	if list.head == nil {
		list.head = newNode
	}

	if list.tail != nil {
		list.tail.SetNextNode(newNode)
		list.tail = newNode
	}
}