package euclid

import "fmt"

func (p *Polynomial) String() string {
	s := ""
	for i := len(*p.coeffs) - 1; i >= 0; i-- {
		if i != 0 {
			c := (*p.coeffs)[i]
			if c != 0 {
				s += fmt.Sprintf("%dx^%d + ", (*p.coeffs)[i], i)
			}
		} else {
			s += fmt.Sprintf("%d", (*p.coeffs)[i])
		}

		if i != 0 {
		}
	}
	s += fmt.Sprintf(" (mod %d)", int(*p.field))
	return s
}

type Polynomial struct {
	coeffs *[]int
	field  *GaloisField
}

func (p *Polynomial) Degree() int {
	return len(*p.coeffs) - 1
}

func PolynomialOver(f *GaloisField) func(...int) *Polynomial {
	return func(coeffs ...int) *Polynomial {
		for i, v := range coeffs {
			coeffs[i] = f.Normalize(v)
		}

		return &Polynomial{
			field:  f,
			coeffs: &coeffs,
		}
	}
}

func sort(p1, p2 *Polynomial) (smaller, larger *Polynomial) {
	if p1.Degree() > p2.Degree() {
		larger = p1
		smaller = p2
	} else {
		larger = p2
		smaller = p1
	}
	return
}

// Add adds two polynomials over a GaloisField
func Add(p1, p2 *Polynomial) *Polynomial {
	if *p1.field != *p2.field {
		panic("must be over same field")
	}

	smaller, larger := sort(p1, p2)

	newCoeffs := make([]int, larger.Degree()+1)

	for i := range *larger.coeffs {
		if i >= len(*smaller.coeffs) {
			newCoeffs[i] = (*(larger.coeffs))[i]
		} else {
			newCoeffs[i] = larger.field.Normalize((*(larger.coeffs))[i] + (*(smaller.coeffs))[i])
		}
	}

	return PolynomialOver(p1.field)(newCoeffs...)
}

// Mul multiplies two polynomials over a GaloisField
func Mul(p1, p2 *Polynomial) *Polynomial {
	if *p1.field != *p2.field {
		panic("must be over the same field")
	}

	smaller, larger := sort(p1, p2)

	newCoeffs := make([]int, larger.Degree()+smaller.Degree()+1)

	for i := range *larger.coeffs {
		c := (*larger.coeffs)[i]
		for j := range *smaller.coeffs {
			degree := i + j
			newCoeffs[degree] = larger.field.Normalize(newCoeffs[degree] + c*(*larger.coeffs)[j])
		}
	}

	return PolynomialOver(p1.field)(newCoeffs...)

}
