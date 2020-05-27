package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	const size = 1000
	var counter int64 = 0
	var wt sync.WaitGroup
	wt.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&counter, 1)
			log.Println("counters", atomic.LoadInt64(&counter))
		}()
		wt.Done()
	}
	wt.Wait()
	log.Println("number of routine", runtime.NumGoroutine())
	log.Println("counter", counter)
}
