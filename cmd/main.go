package main

import (
	"github.com/joelsalt/golang-algorithms/pkg/structure"
)

func main() {
	list := structure.NewSortedLinkedList()
	list.Insert(1).
		Insert(4).
		Insert(3).
		Insert(2).
		Insert(5).
		Insert(5).
		Insert(4).
		Insert(3).
		Insert(2).
		Insert(10)

	list.PrintData()
}