package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	PROD := a * b
	fmt.Printf("PROD = %d\n", PROD)
}
