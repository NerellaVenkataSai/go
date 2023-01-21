// diagram 03 folder 003
// oo in go approach
// go is not oo
// receiver function 009,009.1

package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, card := range cardValues {
			cards = append(cards, card+"of"+suit)
		}
	}

	return cards
}

// reciever function
// print works as a function for type deck
func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}
