package main

import "fmt"

func main() {
	Francisco := 1.50
	Sara := 1.10
	ano := 0

	for Sara <= Francisco {
		Francisco += 0.02
		Sara += 0.03
		ano++
	}
	fmt.Printf("Francisco vai ser mais alto que Sara em %d anos. \n", ano)
}
