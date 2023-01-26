// diagram 03 folder 003
// oo in go approach
// go is not oo
// receiver function 009,009.1

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// reciever function
// print works as a function for type deck
func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

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

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func writeToFile(d deck, fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) deck {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error Read From File:", err)
		os.Exit(1)
	}
	return strings.Split(string(content), ",")
}
