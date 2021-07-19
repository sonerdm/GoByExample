package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

//Sorting our list of people is then as easy as casting the list into our new type. We could also sort by age by doing this:
type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
		{"Soner", 27},
		{"Test", 12},
	}
	//The sort package contains functions for sorting arbitrary data. There are several predefined sorting functions (for slices of ints and floats)
	sort.Sort(ByName(kids))
	fmt.Println("Sorted by name:", kids)
	sort.Sort(ByAge(kids))
	fmt.Println("Sorted by age:", kids)
}
