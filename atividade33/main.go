package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	maiorAB := (a + b + abs(a-b)) / 2
	maior := (maiorAB + c + abs(maiorAB-c)) / 2

	fmt.Printf("%d eh o maior\n", maior)
}
