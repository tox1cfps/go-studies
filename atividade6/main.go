package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Print("Insira a temperatura (Fahrenheit): ")
	fmt.Scan(&fahrenheit)
	celsius := (5 * (fahrenheit - 32) / 9)
	fmt.Printf("Fº: %.2f, Cº: %.2f", fahrenheit, celsius)
}
