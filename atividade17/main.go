package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Insira o seu número: ")
	fmt.Scan(&numero)

	var resultado int
	var contador int

	for {
		if resultado == 1 {
			break
		} else if numero % 2 == 0{
			contador += 1
			resultado = numero / 2
			fmt.Printf("\nO resultado é: %d", resultado)
			numero = resultado
		} else if numero % 2 != 0 {
			contador += 1
			resultado = numero * 3 + 1
			fmt.Printf("\nO resultado é: %d", resultado)
			numero = resultado
		}
	}
		fmt.Printf("\nO algoritimo levou %d passos para chegar em 1", contador)
}