package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wait sync.WaitGroup
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
	}
	wait.Done() 
}
func foo(){
	for i := 0 ; i < 1000 ; i++ {
		fmt.Println("foo:",i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wait.Done()
}