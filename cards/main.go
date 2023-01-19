// package main creates executable file which can be compile and executed using go run or go build
package main

import "fmt"

func main() {

	// var card string = "Ace OF Spades"
	// this is hort syntax for above line which is prefered
	card := "Ace OF Spades"
	card = "override card"

	fmt.Println(card)
}
