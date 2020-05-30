package main

import (
	"fmt"
	"os"
)

var f, err = os.Open("aditya")

func init() {

}
func main() {
	test()
	fmt.Println("hi")
}

func fileClose(f *os.File) {

	fmt.Println("closing file pointer")

	f.Close()
	s := recover()
	fmt.Println("s", s)
}

func test() {
	defer fileClose(f)
	if err != nil {
		panic("adita")
	}
	fmt.Println("sai")
}
