//A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. Functions (also known as procedures or subroutines) are often represented as a black box: (the black box represents the function)

package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

var x int = 5

func main() {
	xs := []float64{98, 93, 77, 82, 82}
	fmt.Println(average(xs))
	f()
	fmt.Println(f1()) //Functions are built up in a “stack”.

	x, y := multiple()
	fmt.Println(x, y)

	fmt.Println(add(1, 2, 3, 4))

	ys := []int{1, 2, 3, 4}
	//We can also pass a slice of ints by following the slice with ...:
	fmt.Println(add(ys...))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(5))

	defer second() //Go has a special statement called defer which schedules a function call to be run after the function completes.
	first()

	wowPanic2()

	fmt.Println("-----Problems------")
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sum of slice elements:", sum(slice))
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(7))
	list := []int{1, 123, 234, 45, 45, 6}
	fmt.Println(greatest(list...))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("fibonacci:", fib(8))
}

//Each time we call a function we push it onto the call stack and each time we return from a function we pop the last function off of the stack.
func f() {
	fmt.Println(x)
}

func f1() int {
	return f2()
}

func f2() (r int) { //We can also name the return type
	r = 3
	return
}

func multiple() (int, int) {
	//Multiple values are often used to return an error value along with the result (x, err := f()), or a boolean to indicate success (x, ok := f()).
	return 5, 6
}

func add(args ...int) int {
	//By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters. In this case we take zero or more ints. We invoke the function like any other function except we can pass as many ints as we want.
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	//makeEvenGenerator returns a function which generates even numbers. Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	//factorial calls itself, which is what makes this function recursive
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

// defer is often used when resources need to be freed in some way. For example when we open a file we need to make sure to close it later. With defer:
/*
f, _ := os.Open(filename)
defer f.Close()
*/
// This has 3 advantages: (1) it keeps our Close call near our Open call so it's easier to understand, (2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and (3) deferred functions are run even if a run-time panic occurs.
func second() {
	fmt.Println("2nd")
}

func wowPanic() {
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}

//But the call to recover will never happen in this case because the call to panic immediately stops execution of the function. Instead we have to pair it with defer:

func wowPanic2() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func greatest(args ...int) int {
	max := args[0]

	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}
