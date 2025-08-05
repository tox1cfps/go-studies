package main

import "fmt"

func main() {
	var produto1, produto2 int
	var unit1, unit2 int
	var price1, price2 float64

	fmt.Scanln(&produto1, &unit1, &price1)
	fmt.Scanln(&produto2, &unit2, &price2)

	calc1 := price1 * float64(unit1)
	calc2 := price2 * float64(unit2)
	total := calc1 + calc2

	fmt.Printf("VALOR A PAGAR: R$ %.02f\n", total)
}
