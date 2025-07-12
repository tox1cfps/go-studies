package main

import "fmt"

func main() {
	var a int
	fmt.Print("Insira o primeiro valor: ")
	fmt.Scan(&a)
	var b int
	fmt.Print("Insira o segundo valor: ")
	fmt.Scan(&b)
	var c int
	fmt.Print("Insira o terceiro valor: ")
	fmt.Scan(&c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("É um triangulo")
	} else {
		fmt.Println("Não é um triangulo")
		return
	}

	if a == b && b == c {
		fmt.Println("É um triangulo Equilatero")
	} else if a == b || b == c || a == c {
		fmt.Println("É um triangulo isósceles ")
	} else {
		fmt.Println("É um triangulo escaleno")
	}
}
