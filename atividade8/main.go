package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Tabuada do %d: \n", i)
		for a := 1; a <= 10; a++ {
			fmt.Printf("%d x %d = %d\n", i, a, i*a)
		}
		fmt.Println()
	}
}
