package main

import (
	"GitHub/GoByExample/An_Introduction_to_Programming_in_Go/chapter11/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
