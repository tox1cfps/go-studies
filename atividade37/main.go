package main

import "fmt"

func main() {
	var valor int
	fmt.Scan(&valor)
	fmt.Println(valor)

	notas100 := valor / 100
	resto := valor % 100

	notas50 := resto / 50
	resto = resto % 50

	notas20 := resto / 20
	resto = resto % 20

	notas10 := resto / 10
	resto = resto % 10

	notas5 := resto / 5
	resto = resto % 5

	notas2 := resto / 2
	resto = resto % 2

	notas1 := resto / 1

	fmt.Printf("%d nota(s) de R$ 100,00\n", notas100)
	fmt.Printf("%d nota(s) de R$ 50,00\n", notas50)
	fmt.Printf("%d nota(s) de R$ 20,00\n", notas20)
	fmt.Printf("%d nota(s) de R$ 10,00\n", notas10)
	fmt.Printf("%d nota(s) de R$ 5,00\n", notas5)
	fmt.Printf("%d nota(s) de R$ 2,00\n", notas2)
	fmt.Printf("%d nota(s) de R$ 1,00\n", notas1)
}
