package main

import (
	"log"
	"runtime"
)

func init(){
	log.Println("Arc",runtime.GOARCH)
	log.Println("cpu",runtime.NumCPU())
	log.Println("os",runtime.GOOS)
	log.Println("Goroutines",runtime.NumGoroutine())
}

func main()  {
	go foo()
	bar()
	log.Println("Goroutines",runtime.NumGoroutine())
}

func foo(){
	for i := 0; i < 1000; i++ {
		log.Println("Goroutines",runtime.NumGoroutine())
		log.Println("foo",i)
	}
}
func bar(){
	for i := 0; i < 1000; i++ {
		log.Println("BAR",i)
	}
}
