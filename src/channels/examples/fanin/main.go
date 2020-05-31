package main

import "fmt"

func main() {
	c := source(1, 2, 3, 4, 5)
	c1 := sqr(c)
	c2 := sqr(c)
	for val := range merge(c1, c2) {
		fmt.Println("val", val)
	}

}

func merge(c1 chan int, c2 chan int) chan int {
	final := make(chan int)
	bol := make(chan bool)
	go func() {
		for c11 := range c1 {
			final <- c11
		}
		bol <- true
	}()
	go func() {
		for c22 := range c2 {
			final <- c22
		}
		bol <- true
	}()
	go func() {
		<-bol
		<-bol
		close(final)
	}()
	return final
}
func source(num ...int) chan int {
	ch := make(chan int)
	go func() {
		for n := range num {
			ch <- n
		}
		close(ch)
	}()
	return ch
}
func sqr(c chan int) chan int {
	out := make(chan int)
	go func() {
		for ch := range c {
			out <- ch * ch
		}
		close(out)
	}()
	return out
}
