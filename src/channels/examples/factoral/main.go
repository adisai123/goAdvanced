package main

import "fmt"

func main() {
	c := factoral(6)
	fmt.Println("factorial of number:", <-c)
}
func factoral(num int) chan int {
	ch := make(chan int)
	fact := 1
	go func() {
		for i := 1; i <= num; i++ {
			fact = fact * i
		}
		ch <- fact
		close(ch)
	}()
	return ch
}
