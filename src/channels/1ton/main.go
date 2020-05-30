package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10000000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for c := range ch {
				fmt.Println(c)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

}
