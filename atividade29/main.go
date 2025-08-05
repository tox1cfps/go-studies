package main

import "fmt"

func main() {
	var name string
	var salary, sold float64

	fmt.Scan(&name, &salary, &sold)

	bonus := sold * 0.15
	total := salary + bonus
	fmt.Printf("TOTAL = R$ %.02f\n", total)
}
