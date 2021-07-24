package main

import (
	"flag"
	"fmt"
)
import "log"
import "example.com/greetings"

func hello2() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}



}

var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help")
}


