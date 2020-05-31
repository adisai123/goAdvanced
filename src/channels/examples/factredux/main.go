package main

import "fmt"

func main() {
	for val := range factoral(generateNumber()) {
		fmt.Println(val)
	}
}

func generateNumber() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				ch <- j
			}
		}
		close(ch)
	}()
	return ch
}
func factoral(c chan int) chan int {
	out := make(chan int)
	go func() {
		for mychan := range c {

			out <- fact(mychan)

		}
		close(out)
	}()
	return out
}

func fact(num int) int {
	fact := 1
	for i := 1; i <= num; i++ {
		fact = fact * i
	}
	return fact
}
