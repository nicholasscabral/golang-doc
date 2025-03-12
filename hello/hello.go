package main

import (
	"fmt"

	"log"

	"rsc.io/quote"

	"example.com/greetings"

	"example.com/asdf"
)

func main() {
	// QUOTE
	fmt.Println(quote.Go())

	// GREETINGS
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Nicholas")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Nicholas", "Cabral", "Dos Anjos"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	for key, value := range messages {
		fmt.Println(key, value)
	}

	// ASDF TEST MODULE
	asdf.PrintRandomDate()
}