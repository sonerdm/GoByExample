package main

import "fmt"

func main() {
	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	fmt.Println(b)

	n2 := 0x9F4
	fmt.Printf("Decimal : %d\n", n2)
	fmt.Printf("Binary : %b\n", n2)
}
