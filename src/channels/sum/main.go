package main

import "fmt"

func main()  {
	
	single()
	double()
}

func single(){
	c := incrementor()
	for ch := range pooler(c){
		fmt.Println(ch)
	}
}
func double(){
	c1 := incrementor()
	c2 := incrementor()
	sum1 := pooler(c1)
	sum2 := pooler(c2)
	fmt.Println( "final counter ",<-sum1 + <-sum2)
	
}

func incrementor() <- chan int {
	ch := make(chan int)
	go func ()  {
		for i := 0; i < 10000000; i++ {
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