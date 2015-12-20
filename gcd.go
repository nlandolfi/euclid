package euclid

import "log"

//  assume a ≥ y ≥ 0, a > 0
func GCD(x, y int) int {
	// Assert that x is less than y
	if x < y {
		return GCD(y, x)
	}

	if y == 0 {
		return x
	}

	return GCD(y, x%y)
}

func EGCD(x, y int) (int, int, int) {
	// Assert that x is less than y
	if x < y {
		return EGCD(y, x)
	}

	if y == 0 {
		return x, 1, 0
	}

	log.Printf("%d = %d x %d + %d", x, x/y, y, x%y)

	d, a, b := EGCD(y, x%y)

	return d, b, (a - (x/y)*b)
}
