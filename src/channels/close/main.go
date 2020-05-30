package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i

			}
			close(c)
		}()
	}

	for c1 := range c {
		fmt.Println(c1)
	}
}
