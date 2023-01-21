// diagram 03 folder
// package main creates executable file which can be compile and executed using go run or go build
package main

import "fmt"

func mainBasic() {

	// var card string = "Ace OF Spades"
	// this is hort syntax for above line which is prefered
	card := "Ace OF Spades"
	card = "override card"

	newCard := newCardF()

	fmt.Println(card)
	fmt.Println(newCard)

	//slice concept
	cards := []string{"Ace OF Spades", newCardF()}
	cards = append(cards, "Some Card")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCardF() string {
	return "Hearts Of Love"
}
