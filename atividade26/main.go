package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	pesoA := 2.0
	pesoB := 3.0
	pesoC := 5.0
	r := a * pesoA
	r2 := b * pesoB
	r3 := c * pesoC
	soma := pesoA + pesoB + pesoC
	media := (r + r2 + r3) / soma
	fmt.Printf("MEDIA = %.1f", media)
}
