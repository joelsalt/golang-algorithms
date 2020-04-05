package algorithm

// InsertionSort is a golang implementation of insertion sort
func InsertionSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		current := data[i] // the current element being inserted
		m := i - 1
		for m >= 0 && data[m] > current {
			data[m + 1] = data[m] // Shift data to the left
			m = m - 1 // Iterate backwards over the array
		}
		data[m + 1] = current // When the while loop is finished, insert the element
	}

	return data
}