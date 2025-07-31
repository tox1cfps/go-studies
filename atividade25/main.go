package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	pesoA := 3.5
	pesoB := 7.5
	r := a * pesoA
	r2 := b * pesoB
	soma := pesoA + pesoB
	media := (r + r2) / soma
	fmt.Printf("MEDIA = %.05f\n", media)
}
