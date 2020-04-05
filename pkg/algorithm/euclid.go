package algorithm

// Euclid is a golang implementation of euclid's algorithm to resolve GCD
func Euclid(a int, b int) int {
	for b != a {
		if b > a {
			b = b - a
		}

		if a > b {
			a = a - b
		}
	}

	return a; // or b
}

// EuclidRecursive is a recursive approach to implementing Euclid's algorithm
func EuclidRecursive(a int, b int) int {
	if b == 0 {
		return a
	}

	return EuclidRecursive(b, a % b)
}

// EuclidArray takes an array of intergers and returns the GCD of all the contained values
func EuclidArray(values []int) int {
	if len(values) == 1 {
		return values[0] // Short-circuit if the array has one element
	} else if len(values) == 2 {
		// Short-circuit if there are two elements, there's no need to loop
		return Euclid(values[0], values[1]) // or EuclidRecursive
	}

	result := Euclid(values[0], values[1])

	// Starting from the third element, compare the rest of the array elements with the result
	for i := 2; i < len(values); i++ {
		result = Euclid(result, values[i])
	}

	return result
}