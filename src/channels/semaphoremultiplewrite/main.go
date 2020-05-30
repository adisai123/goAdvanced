package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 100; i < 200; i++ {
			c <- i
		}
		done <- true
	}()

	go func() { // added it in go routines as range should be ready to accept then only above go routines can put value into it.
		<-done
		<-done
		close(c)
	}()

	for ch := range c {
		fmt.Println(ch)
	}

}
