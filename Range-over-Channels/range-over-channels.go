package main

import "fmt"

func main() {
	//We’ll iterate over 2 values in the queue channel.
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)
	//This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}

//This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
