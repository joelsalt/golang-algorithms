package main

import (
	"github.com/joelsalt/algorithm-implementations/pkg/algorithm"
)

func main() {
	data := []int{5, 4, 1, 3, 2}
	result := algorithm.InsertionSort(data)

	for i := 0; i < len(data); i++ {
		print("Result is: ")
		println(result[i])
	}
}