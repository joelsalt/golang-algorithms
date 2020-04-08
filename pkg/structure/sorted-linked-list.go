package structure

// SortedLinkedList is a singlely-linked list that sorts on insertion
type SortedLinkedList struct {
	head *Node
}

// Insert is a sorted insert into a list
// O(n)
func (list *SortedLinkedList) Insert(data int) *SortedLinkedList {
	newNode := NewNode(data, nil)

	// If either no nodes in the list or the head node is larger than `data`, insert at head
	if list.head == nil || list.head.data > data { 
		return list.headInsert(newNode)
	}

	current := list.head
	for current.nextNode != nil {
		// If the data is larger than current node, and smaller than the next, insert here
		// The >= and < mean that new elements in a sequence will be added at the end of the sequence
		// New element: 2'
		// Sequence: [2, 2, 2, 2, 3]
		// After Sorted Insert => [2, 2, 2, 2, 2', 3]
		if data >= current.data && data < current.nextNode.data {
			newNode.nextNode = current.nextNode
			current.nextNode = newNode
			return list
		}

		current = current.nextNode
	}

	
	// We've reached the end of the list, add the node to the end
	current.nextNode = newNode
	return list
}

func (list *SortedLinkedList) headInsert(node *Node) *SortedLinkedList {
	node.SetNextNode(list.head)
	list.head = node
	return list
}

// NewSortedLinkedList is a factory function for creating a new linked list
func NewSortedLinkedList() SortedLinkedList {
	list := SortedLinkedList{head: nil}
	return list
}

// PrintData loops through and prints all the data in the list
func (list *SortedLinkedList) PrintData() {
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
