package main

import "fmt"

func main() {
	c := make(chan int, 2) //send only type of channnel
	// receive only channel
	c <- 100
	c <- 200

	//m = (chan int) c

	fmt.Print("", <-foo(c)) // not possible
	fmt.Print("", <-foo(c)) // not possible
}

func foo(c <-chan int) <-chan int {
	return c
}
