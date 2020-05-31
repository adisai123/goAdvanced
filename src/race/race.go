package main

import (
	"log"
	"os"
	"sync"
)

var wait sync.WaitGroup
var counter int = 0

func main() {
	fs, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(fs)
	for rout := 1; rout <= 2; rout++ {

		wait.Add(1)
		go routine(rout)
	}

	wait.Wait()
	log.Printf("Final counter: %d\n", counter)
}

func routine(id int) {

	for count := 0; count < 2; count++ {

		value := counter
		value++
		counter = value
	}

	wait.Done()
}
