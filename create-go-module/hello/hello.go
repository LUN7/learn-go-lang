package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
    // Get a greeting message and print it.
		log.SetPrefix("greetings: ")
		log.SetFlags(0)
    message, error := greetings.Hello("Gladys")
		if error!=nil {
			log.Fatal(error)
		}
    fmt.Println(message)
}
