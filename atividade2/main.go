package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Insira seu nome: ")
	fmt.Scan(&nome)
	var idade int
	fmt.Print("Insira sua idade: ")
	fmt.Scan(&idade)

	if idade < 18 {
		fmt.Printf("%s É menor de idade\n", nome)
	} else {
		fmt.Printf("%s É maior de idade\n", nome)
	}
}
