package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			fmt.Print(i, " even ")
		} else {
			fmt.Print(i, " odd ")
		}

		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")

		}
	}
}
