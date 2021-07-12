//Making progress on more than one task simultaneously is known as concurrency. Go has rich support for conc

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) { //channel direction
	for i := 0; ; i++ {
		c <- "ping"
		// The <- (left arrow) operator is used to send and receive messages on the channel. c <- "ping" means send "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c //means receive a message and store it in msg
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {

	// This program consists of two goroutines. The first goroutine is implicit and is the main function itself.
	//The second goroutine is created when we call go f(0). Normally when we invoke a function our program will
	//execute all the statements in a function and then return to the next line following the invocation. With
	//a goroutine we return immediately to the next line and don't wait for the function to complete. This is why
	//the call to the Scanln function has been included; without it the program would exit before being given the
	//opportunity to print all the numbers.

	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// 	//Goroutines are lightweight and we can easily create thousands of them. We can modify our program to run 10 goroutines by doing this:
	// }
	// var input string
	// fmt.Scanln(&input)
	fmt.Println("-----Channels------")
	// Channels provide a way for two goroutines to communicate with one another and synchronize their execution.

	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	var input string
	// fmt.Scanln(&input)

	fmt.Println("-----Select------")
	//Go has a special statement called select which works like a switch but for channels

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready")
			}
		}
	}()
	fmt.Scanln(&input)

	fmt.Println("-----Buffered Channels------")
	//c := make(chan int, 1)
	// It's also possible to pass a second parameter to the make function when creating a channel:
	// This creates a buffered channel with a capacity of 1. Normally channels are synchronous; both
	// sides of the channel will wait until the other side is ready. A buffered channel is asynchronous;
	// sending or receiving a message will not wait unless the channel is already full.

	fmt.Println("-----Problems------")

	// 	1. How do you specify the direction of a channel type?

	// You can specify the direction of a channel using the function signature and the arrow signs:
	// c chan<- string means that we can only send to c
	// c <-chan string means that we can only receive to c

	fmt.Println(time.Now())
	sleep(3)
	fmt.Println(time.Now())
}

func sleep(secs int) {
	select {
	case <-time.After(time.Duration(secs) * time.Second):
		return
	}
}

// 3. What is a buffered channel? How would you create on with a capacity of 20?
// A buffered channel is an asynchronous channel that does not wait unless is at full capacity.
// c := make(chan int, 20)
