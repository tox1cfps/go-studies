package main

import "fmt"

func main() {
	var r float64
	fmt.Scan(&r)
	pi := 3.14159

	sphere := (4.0 / 3) * pi * r * r * r
	fmt.Printf("VOLUME = %.03f\n", sphere)
}
