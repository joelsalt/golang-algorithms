package main

import (
	"github.com/joelsalt/golang-algorithms/pkg/structure"
	"github.com/joelsalt/golang-algorithms/pkg/algorithm"
)

func main() {
	list := structure.NewDoublyLinkedList()
	list.Insert(1).
		Insert(4).
		Insert(1).
		Insert(3).
		Insert(2).
		Insert(5).
		Insert(5).
		Insert(4).
		Insert(3).
		Insert(2).
		Insert(10)

	sortedList := algorithm.LinkedListInsertionSort(list)

	sortedList.PrintData()
}