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
	fmt.Printf("Sara vai ser mais alta que Francisco em %d anos. \n", ano)
}
