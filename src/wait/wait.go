package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wait sync.WaitGroup
var counter int 
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main(){
	wait.Add(2)
	go foo()
	go bar()
	wait.Wait()
	

}

func bar(){
	for i := 0 ; i < 1000 ; i++ {
		fmt.Println("bar:",i)
		time.Sleep(time.Duration(3 * time.Millisecond))
		counter ++
		printCounter("bar counter:",counter)
	}

	wait.Done() 
}
func foo(){
	for i := 0 ; i < 1000 ; i++ {
		fmt.Println("foo:",i)
		time.Sleep(time.Duration(3 * time.Millisecond))
		printCounter("foo counter:",counter)
	}


	wait.Done()
}
func printCounter(msg string , i int){
	fmt.Println(msg,counter)
}