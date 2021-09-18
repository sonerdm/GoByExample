package main

import "fmt"

func main() {
	n := 2548
	fmt.Printf("%x\n", n) //You can also use"%X" to print a hexadecimal number with capitalized letters :
	n2 := 2548
	n3 := 0x9F4
	fmt.Printf("%X\n", n2)
	fmt.Printf("%x\n", n3)

	n4 := 0x9F4
	fmt.Printf("Decimal : %d\n", n4)
	fmt.Printf("Octal : %o\n", n4)
	fmt.Printf("Octal with 0o prefix: %O\n", n4)

	//"%o" allow you to print the number in octal
	//"%O" allow you to print the number in octal with a"0o" prefix
}
