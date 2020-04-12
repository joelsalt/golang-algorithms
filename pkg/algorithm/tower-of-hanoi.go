package algorithm

import(
	"strconv"
)

// TowerOfHanoi moves all n elements across the pegs and prints the steps
func TowerOfHanoi(numberOfDiscs int) {
	move(numberOfDiscs, "A", "C", "B")
}

func move (n int, source string, target string, spare string) {
	if n == 1 {
		println( "Moving disc 1 from " + source + " to " + target )
	} else {
		move(n - 1, source, spare, target)
		println( "Moving disc " + strconv.Itoa(n) + " from " + source + " to " + target )
		move(n - 1, spare, target, source)
	}
}