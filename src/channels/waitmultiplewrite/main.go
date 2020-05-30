package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wt sync.WaitGroup
	wt.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		wt.Done()
	}()
	go func() {
		for i := 100; i < 200; i++ {
			c <- i
		}
		wt.Done()
	}()
	go func() {
		wt.Wait()
		close(c)
	}()
	for ch := range c {
		fmt.Println(ch)
	}
}
