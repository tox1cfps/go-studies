package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Print("\nInsira o número A: ")
	fmt.Scan(&a)
	fmt.Print("\nInsira o número B: ")
	fmt.Scan(&b)
	fmt.Print("\nInsira o número C: ")
	fmt.Scan(&c)

	if a > b && a > c {
		fmt.Printf("%d eh o maior", a)
	} else if b > a && b > c {
		fmt.Printf("%d eh o maior", b)
	} else {
		fmt.Printf("%d eh o maior", c)
	}
}
