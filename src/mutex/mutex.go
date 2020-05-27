package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var counter int
var wait sync.WaitGroup

func main() {
	wait.Add(2)
	go incrementFoo("aditya")
	go incrementFoo("nupur")
	wait.Wait()
	fmt.Println("final count", counter)
}

func incrementFoo(incrementor string) {
	for i := 0; i < 50; i++ {
		mutex.Lock()
		time.Sleep(4 * time.Second)
		counter++
		fmt.Println(incrementor, i, counter)
		mutex.Unlock()
	}
	wait.Done()
}
