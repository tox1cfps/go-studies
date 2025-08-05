package main

import "fmt"

func main() {
	var number, hours, hourSalary float64
	fmt.Scan(&number, &hours, &hourSalary)

	var salary float64

	salary = hours * hourSalary
	fmt.Printf("NUMBER = %.f\nSALARY = U$ %.02f\n", number, salary)
}
