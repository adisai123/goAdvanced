package main

import "fmt"

func main() {
	c := make(chan int, 10000)
	c <- 100 // channnels are block , it can not execute next line as its blocked unless buffer specified
	for i := 0; i < 10000; i++ {
		go increment(c, i)
	}
	fmt.Println(<-c) // send and receive can not occur at the same time.
}

func increment(c chan int, i int) {
	c <- i
}
