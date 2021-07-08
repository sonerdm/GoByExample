package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("-----Arrays------")
	var x [5]int //definition of array
	x[4] = 100
	fmt.Println("first array", x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	var total2 float64 = 0
	for i := 0; i < len(y); i++ {
		total2 += y[i]
	}
	fmt.Println(total2 / float64(len(y)))
	//The Go compiler won't allow you to create variables that you never use. Since we don't use i inside of our loop we need to change it to this:
	//A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)
	var total3 float64 = 0
	for _, value := range y {
		total3 += value
	}
	fmt.Println(total / float64(len(y)))

	z := [5]float64{98, 93, 77, 82, 83} // shorter syntax for creating arrays:

	z2 := [5]float64{ //Notice the extra trailing , after 83. This is required by Go and it allows us to easily remove an element from the array by commenting out the line:
		98,
		//93,
		77,
		82,
		//83,
	}
	fmt.Println(z)
	fmt.Println(z2)

	fmt.Println("-----Slices------")
	//A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change.

	//var x2 []float64 //an example of a slice

	//If you want to create a slice you should use the built-in make function:
	//x2 := make([]float64, 5)

	// 10 represents the capacity of the underlying array which the slice points to:
	//x2 := make([]float64, 5, 10)

	// For convenience we are also allowed to omit low, high or even both low and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5] and arr[:] is the same as arr[0:len(arr)]
	sliceAppend()
	sliceCopy()

	fmt.Println("-----Maps------")

	//A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key. Here's an example of a map in Go
	//var x map[string]int
	myFirstMap()
	mapExample()
	advancedMap()

	fmt.Println("-----Problems------")
	problems()
}

func sliceAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("append:", slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("copy:", slice1, slice2)
}

func myFirstMap() {
	//maps have to be initialized before they can be used
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	y[2] = 11
	fmt.Println(y)
	delete(y, 2)
	fmt.Println("deleted", y)
}

func mapExample() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	//Suppose we tried to look up an element that doesn't exist, you should see nothing returned.
	//Technically a map returns the zero value for the value type (which for strings is the empty string).
	fmt.Println(elements["Un"])

	//Accessing an element of a map can return two values instead of just one. The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
	// name, ok := elements["F"]
	// fmt.Println(name, ok)

	//First we try to get the value from the map, then if it's successful we run the code inside of the block.
	if name, ok := elements["B"]; ok {
		fmt.Println(name, ok)
	}
}

func advancedMap() {
	//Notice that the type of our map has changed from map[string]string to map[string]map[string]string. We now have a map of strings to maps of strings to strings. The outer map is used as a lookup table based on the element's symbol, while the inner maps are used to store general information about the elements.
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Ne"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func problems() {
	x := [4]int{1, 2, 3, 4}
	fmt.Println("4th element:", x[3])

	slice := make([]int, 3, 9)
	fmt.Println("length:", len(slice), "capacity:", cap(slice))

	array := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(array[2:5])

	list := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sort.Ints(list)
	fmt.Println("smallest number:", list[0])

}

// 6. What are defer, panic and recover? How do you recover from a run-time panic?

// defer schedules a function call to be run after the completion of the function
// panic immediately stops the execution of the function
// recover stops the panic and returns the value that was passed to the call to panic

// We can recover from a run-time panic using the recover built-in.
