package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	pi := 3.14159

	triangulo := a * c / 2
	circulo := pi * c * c
	trapezio := (a + b) * c / 2
	quadrado := b * b
	retangulo := a * b

	fmt.Printf("TRIANGULO: %0.3f\nCIRCULO: %0.3f\nTRAPEZIO: %0.3f\nQUADRADO: %0.3f\nRETANGULO: %.03f\n", triangulo, circulo, trapezio, quadrado, retangulo)
}
