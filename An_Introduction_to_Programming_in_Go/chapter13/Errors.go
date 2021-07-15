package main

//Go has a built-in type for errors that we have already seen (the error type). We can create our own errors by using the New function in the errors package:
import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error message")
	fmt.Println(err)
}
