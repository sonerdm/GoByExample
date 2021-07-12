package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	y := 10
	zero(x)
	zeroPointer(&y)
	//Pointers reference a location in memory where a value is stored rather than the value itself.
	//(They point to something else) By using a pointer (*int) the zeroPointer function is able to modify the original variable.
	fmt.Println(x)
	fmt.Println("pointer:", y)

	xPtr := new(int) //Another way to get a pointer is to use the built-in new function
	//new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	one(xPtr)
	fmt.Println(*xPtr)

	fmt.Println("-----Problems-----")
	// 	1. How do you get the memory address of a variable?
	// We can use the & operator to find the address of a variable.

	// 2. How do you assign a value to a pointer?
	// Using the * operator we can assign a value to a pointer.

	// 3. How do you create a new pointer?
	// Using the new built-in function

	z := 1.5
	square(&z)
	c := 4
	d := 8
	fmt.Println(swap(1, 2))
	swapP(&c, &d)
	fmt.Println(c, d)
}

func one(xPtr *int) {
	*xPtr = 1
}

//In some programming languages there is a significant difference between using new and &, with great care being needed to eventually delete anything
//created with new. Go is not like this, it's a garbage collected programming language which means memory is cleaned up automatically when nothing refers to it anymore.
// Pointers are rarely used with Go's built-in types, but as we will see in the next chapter, they are extremely useful when paired with structs.

func square(z *float64) {
	*z = *z * *z
	fmt.Println(*z)
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapP(c *int, d *int) {
	temp := *c
	*c = *d
	*d = temp
}
