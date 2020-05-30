package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	f := -100.5
	n, err := divideBy2(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("result of number", n)
}

func divideBy2(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Please dont provide negative number")
	}
	return f / 2, nil
}
