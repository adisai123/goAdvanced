package main

import (
	"fmt"
)

func main() {
	c := fanin(boaring("aditya"), boaring("nupur"))
	for i := 0; i < 50; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("you're both boaring, I am leaving")
}

func boaring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s,%d", msg, i)
			//time.Sleep(1 * time.Second)
		}
	}()

	return c
}
func fanin(i1, i2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-i1
		}
	}()

	go func() {
		for {
			c <- <-i2
		}
	}()
	return c
}
