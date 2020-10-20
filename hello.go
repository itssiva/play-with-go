package main

import (
	"fmt"
	"example.com/greetings"
)

// import "rsc.io/quote"

// func main() {
// 		fmt.Println(quote.Go());
// 		fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
// }

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}