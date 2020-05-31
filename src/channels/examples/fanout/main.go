package main

import (
	"fmt"
	"time"
)

var producerID int
var consumerID int

func main() {
	c := make(chan string)
	for {
		go producer(c)
		go producer(c)
		go producer(c)
		go consumer(c)
		go consumer(c)
		go consumer(c)
		time.Sleep(1 * time.Second)
	}
}

func producer(c chan string) {
	producerID++
	thisID := producerID
	dataID := 0
	for {
		fmt.Println("producer prodicing ", thisID)
		dataID++
		c <- fmt.Sprintf("data from publisher %d . data  %d", thisID, dataID)
	}
}

func consumer(c chan string) {
	consumerID++
	for {
		fmt.Printf("%d consuming from %s", consumerID, <-c)
	}
}
