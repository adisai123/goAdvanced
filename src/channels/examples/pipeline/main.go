package main

import "fmt"

func main() {
	for sch := range sqr(sqr(sqr(getChannel(2, 3, 4, 5, 6, 7)))) {
		fmt.Println(sch)
	}
}

func getChannel(nums ...int) chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func sqr(c chan int) chan int {
	ch := make(chan int)
	go func() {
		for cr := range c {
			ch <- cr * cr
		}
		close(ch)
	}()
	return ch
}
