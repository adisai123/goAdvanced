package main

import "fmt"

func main() {
	c := make(chan int) //send only type of channnel
	// receive only channel
	go bar(c)
	foo(c)

	//m = (chan int) c

}

func foo(c <-chan int) {
	//range over a channel ; blocks until channel close
	for item := range c {
		fmt.Println("aditya", item)
	}
}
func bar(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("adotya")
	}
	close(c)
}
