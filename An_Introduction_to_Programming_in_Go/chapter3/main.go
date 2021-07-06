package main

import "fmt"

func main() {
	fmt.Println("-----Numbers-----")
	fmt.Println("1 + 1 =", 1.0+1.0) //addition
	fmt.Println("1 - 1 =", 1.0-1.0) //subtraction
	fmt.Println("1 * 1 =", 1.0*1.0) //multiplication
	fmt.Println("1 / 1 =", 1.0/1.0) //division
	fmt.Println("1 % 1 =", 1%1)     //remainder

	fmt.Println("-----Strings-----")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println("-----Booleans-----")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
