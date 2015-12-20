package euclid

type GaloisField int

func (gf *GaloisField) Normalize(a int) int {
	for a < 0 {
		a += int(*gf)
	}

	return a % int(*gf)
}

func (gf *GaloisField) Inverse(a int) int {
	return gf.Normalize(Inverse(a, int(*gf)))
}

// the multiplicative inverse of a mod p
// assuming a and p are coprime
// assume a < p
func Inverse(a, p int) int {
	_, _, y := EGCD(p, a)
	return y
}
