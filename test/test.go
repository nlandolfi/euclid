package main

import (
	"fmt"
	"log"

	"github.com/nlandolfi/euclid"
)

func main() {
	log.Printf("GCD(872, 208) = %d", euclid.GCD(872, 208))
	log.Printf("GCD(1952, 872) = %d", euclid.GCD(872, 208))

	d, a, b := euclid.EGCD(2328, 440)
	log.Printf("EGCD(2338, 440); d = %d, a = %d, b = %d", d, a, b)

	gf := euclid.GaloisField(7)
	log.Printf("Inverse(6, 7): %d", gf.Inverse(6))

	p := euclid.PolynomialOver(&gf)(1, 2, 3)
	p2 := euclid.PolynomialOver(&gf)(4, 5, 6)
	fmt.Printf("\n   %s\n + %s\n", p, p2)
	fmt.Println("----------------------------")
	fmt.Println("  ", euclid.Add(p, p2))

	plus := euclid.PolynomialOver(&gf)(1, 1)
	minus := euclid.PolynomialOver(&gf)(-1, 1)
	fmt.Printf("\n   %s\n * %s \n", plus, minus)
	fmt.Println("----------------------------")
	fmt.Println("  ", euclid.Mul(plus, minus))
}
