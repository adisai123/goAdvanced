package main

import (
	"log"
	"runtime"
	"sync"
)

func init(){
	log.Println("Arc",runtime.GOARCH)
	log.Println("cpu",runtime.NumCPU())
	log.Println("os",runtime.GOOS)
	log.Println("Goroutines",runtime.NumGoroutine())
}
//code contains race conditions
func main()  {
	go foo()
	bar()
	log.Println("Goroutines",runtime.NumGoroutine())
	const size = 1000
	var wt sync.WaitGroup
	wt.Add(size)
	counter := 0 
	// /var mu sync.Mutex
	for i := 0; i < size; i++ {
		go func ()  {
			//mu.Lock()
			mycnt := counter
			runtime.Gosched()
			mycnt ++
			counter = mycnt
			wt.Done()
			//mu.Unlock()
		}()
	}
	wt.Wait()
	log.Println("count",counter)
	log.Println("Goroutines",runtime.NumGoroutine())
}

func foo(){
	for i := 0; i < 1000; i++ {
		log.Println("foo",i)
	}
}
func bar(){
	for i := 0; i < 1000; i++ {
		log.Println("BAR",i)
	}
}
