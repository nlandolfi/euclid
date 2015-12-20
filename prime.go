package euclid

// Trivial (slow) primality (very slow)
func Prime(p int) bool {
	if p < 2 {
		return false
	} else if p == 2 {
		return true
	}

	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}

	return true

}

func Coprime(x, y int) bool {
	return GCD(x, y) == 1
}
