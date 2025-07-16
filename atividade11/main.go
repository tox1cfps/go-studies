package main

import "fmt"

func main() {
	var numero1 int
	fmt.Print("Insira o valor: ")
	fmt.Scan(&numero1)
	var numero2 int
	fmt.Print("Insira o valor: ")
	fmt.Scan(&numero2)

	resultado := numero1 % numero2
	fmt.Println(numero1 / numero2)
	fmt.Println(resultado)
}
