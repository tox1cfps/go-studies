package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite o numero que deseja ver a tabuada: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
