package notes

// 5. SLICES
func sliceCreation() {
	numbers := []int{1, 3, 5, 7, 9}
	println(numbers[0]) // '1'
}

func sliceCreationWithMake() {
	// 'make' allocates memory and initializes the slice
	numbers := make([]int, 10)
	println(numbers[0]) // '0' - slice is initialized with 0s for each element
}

func sliceResizing() {
	numbers := make([]int, 0, 2)
	numbersCap := cap(numbers)
	println(numbersCap) // '2'

	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)

		// if capacity has changed, Go grew the array by x2 to accommodate
		if cap(numbers) != numbersCap {
			numbersCap = cap(numbers)
			println(numbersCap) // '4', then '8' due to the 2x growth
		}
	}
}

func removeElementFromSlice() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = removeAtIndex(numbers, 2)
	println(numbers) // '[1, 2, 4, 5]'
}

func removeAtIndex(sourceSlice []int, index int) []int {
	lastIndex := len(sourceSlice) - 1
	// swap the value to remove with the last value
	sourceSlice[index], sourceSlice[lastIndex] = sourceSlice[lastIndex], sourceSlice[index]
	return sourceSlice[:lastIndex]
}