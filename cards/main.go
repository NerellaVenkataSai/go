package main

func main() {
	// for go basics
	// mainBasic()

	// cards := deck{"Ace Of Spades", "Love Hearts"}
	cards := newDeck()
	hand, remainingCrds := deal(cards, 4)

	cards.print()
	hand.print()
	remainingCrds.print()
	writeToFile(cards, "firstWrite")
	content := readFromFile("firstWrite")
	content.print()

	cards.shuffleCards().print()

	// assignment
	evenOrOdd([]int{1, 2, 3, 4, 5})
}
