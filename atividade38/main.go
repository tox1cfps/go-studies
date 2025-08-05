package main

import "fmt"

func main() {
	var time int
	fmt.Scan(&time)

	hours := time / 3600
	minutes := (time % 3600) / 60
	seconds := time % 60

	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}
