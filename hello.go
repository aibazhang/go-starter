package main

import (
	"fmt"
	"log"

	"github.com/aibazhang/go-starter/greetings"
)

func main() {
	// Set properties of the predefine Logger, include
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message
	message, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was return, print the returned message
	// to the console
	fmt.Print(message)
}
