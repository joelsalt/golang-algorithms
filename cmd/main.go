package main

import (
	"github.com/joelsalt/golang-algorithms/pkg/structure"
)

func main() {
	bTree := new(structure.BinaryTree)
	bTree.Insert(20).
		Insert(10).
		Insert(30).
		Insert(5).
		Insert(15).
		Insert(25).
		Insert(35)
	
	bTree.TraverseInOrder()
}