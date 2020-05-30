package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	for i := 0; i < 100; i++ {
		go func() {
			for i := i; i < i+100; i++ {
				ch <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < 100; i++ {
			<-done

		}
		close(ch)
	}()
	for c := range ch {
		fmt.Println(c)
	}
}
