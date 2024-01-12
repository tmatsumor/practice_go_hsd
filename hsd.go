// package hsd provide functions for calculate Hamming distance.
package hsd

// convert string arguments to rune array.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// calculate Hamming distance.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}

func Hoge() int {
	return 0
}

func Bar() int {
	return 1
}

func Mya() int {
	return 2
}
