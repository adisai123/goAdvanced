package main

import "fmt"

func main()  {
	c := incrementor()
	sum := pooler(c)
	for ch := range sum {
		fmt.Println(ch)
	}

}

func incrementor() <- chan int {
	ch := make(chan int)
	go func ()  {
		for i := 0; i < 100; i++ {
			ch <- i 
		}
		close(ch)
	}()
	return ch 
}

func pooler(c <- chan int) <- chan int{
	out := make(chan int)
	go func() {
		var sum int 
		for ch := range c {
			sum += ch
		}
		out <- sum
		close(out)
	}()
	return out
}