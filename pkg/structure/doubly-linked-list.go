package structure

// DoublyLinkedList is a golang implementation of the data structure
type DoublyLinkedList struct {
	Head *DoublyLinkedNode
}

// Insert adds a DoublyLinkedNode to a doubly linked list
func (list *DoublyLinkedList) Insert(data int) *DoublyLinkedList {
	newNode := NewDoublyLinkedNode(data, nil, nil)
	
	// Start: nil <= HeadNode => ...

	// Set the new node to point to the start of the list
	newNode.NextNode = list.Head //  newNode => HeadNode

	if list.Head != nil {
		// point the Head node's back reference towards the new node
		list.Head.PreviousNode = newNode // newNode <= HeadNode
	}

	// Move the Head to point to the new node at the beginning of the list 
	list.Head = newNode // nil <= newHeadNode => oldHeadNode
	return list
}

// Length return the length of the linked list
func (list *DoublyLinkedList) Length() int {
	counter := 0
	current := list.Head
	for current != nil {
		counter++
		current = current.NextNode
	}

	return counter
}

// NewDoublyLinkedList is a factory function for creating a new linked list
func NewDoublyLinkedList() DoublyLinkedList {
	list := DoublyLinkedList{Head: nil}
	return list
}

// PrintData loops through and prints all the data in the list
func (list *DoublyLinkedList) PrintData() {
	current := list.Head
	counter := 0

	for (current != nil) {
		print("Node element ")
		print(counter)
		print(" has value: ")
		println(current.data)

		counter++
		current = current.NextNode
	}
}

// PrintDataReverse loops through and prints all the data in the list
func (list *DoublyLinkedList) PrintDataReverse() {
	tail := list.Head
	for tail != nil {
		tail = tail.NextNode
	}

	current := tail
	counter := 0

	for (current != nil) {
		print("Node element ")
		print(counter)
		print(" has value: ")
		println(current.data)

		counter++
		current = current.PreviousNode
	}
}
