package notes

// 1. VARIABLES
func variableDefinitions() {
	var power1 int
	var power2 int = 9000
	power3 := false

	// '_' allows you to ignore returned values from a function
	_, _, boolean := multiReturnValues(power1, power2, power3)
	println(boolean)
}

func multiReturnValues(v1 int, v2 int, v3 bool) (int, int, bool) {
	return v1, v2, v3
}
