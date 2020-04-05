package notes

// 4. ARRAYS
func arrayInitialization() {
	var scores [10]int
	scores[0] = 339
	length := len(scores)
	println( length )

	otherScores := [4]int{1, 2, 3, 4}
	println(otherScores[0])
}

func arrayIteration() {
	scores := [4]int{1, 2, 3, 4}

	for index, value := range scores {
		println(index, value)
	}
}
