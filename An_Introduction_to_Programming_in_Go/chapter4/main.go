package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	const x string = "Hello World" //Constants
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("Your number doubled:", output)
	fmt.Println("a:", a, "b", b, "c", c)
	fmt.Println(x)
}
