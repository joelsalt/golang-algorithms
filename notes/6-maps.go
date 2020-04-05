package notes

// 6. MAPS
func mapCreation() {
	// Maps function like hashtables, and are also created with the make function
	saiyans := make(map[string]int)
	saiyans["goku"] = 9001

	power, exists := saiyans["goku"]
	println(power, exists) // '9001, true'
}

func mapCreationWithSize() {
	// Since maps are dynamic, defining size at creation can help with performance
	saiyans := make(map[string]int, 100)
	saiyans["goku"] = 9001

	delete(saiyans, "goku")
}

func mapInitialization() {
	saiyans := map[string]int{
		"goku": 9001,
		"gohan": 2044
	}
}

func mapIteration() {
	saiyans := map[string]int{
		"goku": 9001,
		"gohan": 2044
	}

	// iteration over maps is non-deterministic. The keys/values will be returned in a random order
	for key, value := range saiyans {
		println(key, value) // '"Goku", 9001 /n "Gohan", 2044'
	}
}
