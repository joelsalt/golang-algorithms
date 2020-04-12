package algorithm

import (
	"math"
)

// MergeSort is a golang implementation of the sort algorithm
func MergeSort(data []int, start int, end int) {
	if start < end {
		middle := int( math.Floor((float64) (start + end) / 2) )
		MergeSort(data, start, middle)
		MergeSort(data, middle + 1, end)
		merge(data, start, middle, end)
	}
}

func merge(data []int, start int, middle int, end int) {
	n1 := middle - start + 1
	n2 := end - middle
	left := make([]int, n1)
	right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = data[start + i]
	}
	for j := 0; j < n2; j++ {
		right[j] = data[middle + 1 + j]
	}

	n, m := 0, 0
	for i := start; i <= end; i++ {
		if left[n] <= right[m] {
			data[i] = left[n]
			if (n < len(left) - 1) {
				n = n + 1
			}
		} else {
			data[i] = right[m]
			if(m < len(right) - 1) {
				m = m + 1
			}
		}
	}
}