package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"example.com/greetings"

	"example.com/asdf"
)

func main() {
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Nicholas")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	asdf.PrintRandomDate()
}