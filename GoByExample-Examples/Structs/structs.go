package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	//This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	//You can name the fields when initializing a struct.
	fmt.Println(person{name: "Soner", age: 27})

	//Omitted fields will be zero-valued.
	fmt.Println(person{name: "Tugba"})

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Merve", age: 27})

	//It’s idiomatic to encapsulate new struct creation in
	//constructor functions
	fmt.Println(newPerson("Efe"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
