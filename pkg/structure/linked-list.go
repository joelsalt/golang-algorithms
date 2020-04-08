package structure

// LinkedList is a go implementation of a Linked List
type LinkedList struct {
	head *Node
}

// Insert adds a new node at the beginning of the linked list
// O(1)
func (list *LinkedList) Insert(data int) *LinkedList {
	newNode := NewNode(data, list.head)
	list.head = newNode
	return list
}

// Delete removes the 0th element from the linked list
// O(1)
func (list *LinkedList) Delete() *LinkedList {
	list.head = list.head.nextNode
	return list
}

// Find is a search implementation for linked lists
func (list *LinkedList) Find(data int) *Node {
	current := list.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.nextNode
	}

	return nil
}

// Length return the length of the linked list
func (list *LinkedList) Length() int {
	counter := 0
	current := list.head
	for current != nil {
		counter++
		current = current.nextNode
	}

	return counter
}

// PrintData loops through and prints all the data in the list
func (list *LinkedList) PrintData() {
	current := list.head
	counter := 0

	for (current != nil) {
		print("Node element ")
		print(counter)
		print(" has value: ")
		println(current.data)

		counter++
		current = current.nextNode
	}
}

// NewLinkedList is a factory function for creating a new linked list
func NewLinkedList() LinkedList {
	list := LinkedList{head: nil}
	return list
}
