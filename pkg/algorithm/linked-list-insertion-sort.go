package algorithm

import (
	"github.com/joelsalt/golang-algorithms/pkg/structure"
)

// LinkedListInsertionSort takes a doubly linked list and sorts it via insertion
func LinkedListInsertionSort(list structure.DoublyLinkedList) structure.DoublyLinkedList {
	current := list.Head.NextNode // Skip the first element, as it can't be moved any further
	for current != nil {
		nextNode := current.NextNode // We want to keep a reference to the unsorted next element
														   // in case current is moved
		insertAfter := current.PreviousNode // insertAfter represents the node we'll stop on
		for insertAfter != nil && insertAfter.GetData() > current.GetData() {
			insertAfter = insertAfter.PreviousNode
		}

		// Only run this code when insertAfter is a node or when the current node is not the list head
		if (insertAfter != nil || current.PreviousNode != nil) {
			current.PreviousNode.NextNode = current.NextNode
			if current.NextNode != nil {
				current.NextNode.PreviousNode = current.PreviousNode
			}
			current.PreviousNode = insertAfter
		}
		
		if insertAfter == nil { // Update the head of the list
			current.NextNode = list.Head
			current.NextNode.PreviousNode = current
			current.PreviousNode = nil
			list.Head = current
		} else { // Only update the insertAfter node when it is not nil
			current.NextNode = insertAfter.NextNode
			insertAfter.NextNode.PreviousNode = current
			insertAfter.NextNode = current
		}
		current = nextNode
	}

	return list
}

func insertElement(
	element *structure.DoublyLinkedNode,
	insertAfter *structure.DoublyLinkedNode,
	list structure.DoublyLinkedList,
	) {
	if element.PreviousNode != nil {
		element.PreviousNode.NextNode = element.NextNode
	}
	
	if element.NextNode != nil {
		element.NextNode.PreviousNode = element.PreviousNode
	}

	element.PreviousNode = insertAfter

	
	if insertAfter != nil {
		element.NextNode = insertAfter.NextNode

		insertAfter.NextNode = element
	} else {
		element.NextNode = list.Head

		list.Head.PreviousNode = element

		element.PreviousNode = nil
	
		list.Head = element
	}
}