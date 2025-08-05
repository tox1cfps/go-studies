package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2 float64
	var y1, y2 float64
	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)

	distance := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	fmt.Printf("%.04f\n", distance)
}
