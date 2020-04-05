package algorithm

// SelectionSort is a golang implementation of selection sort
func SelectionSort(data []int) []int {
	for i := 0; i < len(data) - 1; i++ {
		minIndex := i // Find the index of the next smallest element
		for m := i + 1; m < len(data); m++ {
			if data[m] < data[minIndex] {
				minIndex = m
			}
		}

		data[i], data[minIndex] = data[minIndex], data[i] // Swap elements on each pass
	}

	return data
}