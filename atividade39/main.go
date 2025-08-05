package main

import "fmt"

func main() {
	var totalDays int
	fmt.Scan(&totalDays)

	years := totalDays / 365
	resto := totalDays % 365
	months := resto / 30
	days := resto % 30

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", years, months, days)
}
