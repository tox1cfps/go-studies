package main

import "fmt"

func main() {
	var X int
	var Y float64
	fmt.Scan(&X, &Y)

	result := float64(X) / float64(Y)
	fmt.Printf("%.03f km/l\n", result)
}
