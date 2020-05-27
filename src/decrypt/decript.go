package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	log.Println("me init functuin ")
}

func main() {

	s := "nupurchandak"

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))
	comperr := bcrypt.CompareHashAndPassword(bs, []byte("nupurchandaks"))
	if comperr != nil {
		log.Fatal("password is  not same")

	} else {
		log.Println("password is same")
	}

}

func init() {
	log.Println("second func")
}
