package main

import "fmt"

func main() {
	var a string
	fmt.Print("Insira o valor A: ")
	fmt.Scan(&a)
	var b string
	fmt.Print("Insira o valor B: ")
	fmt.Scan(&b)

	temp := a
	a = b
	b = temp

	fmt.Println("Novo valor de A: ", a)
	fmt.Println("Novo valor de B: ", b)

}
