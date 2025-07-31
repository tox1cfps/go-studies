package main

import "fmt"

func main() {
	var R, A float64
	fmt.Scan(&R)
	n := 3.14159
	A = n * R * R
	fmt.Printf("A=%.04f\n", A)
}
