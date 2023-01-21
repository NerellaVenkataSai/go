// diagram 03 folder 003
// oo in go approach
// go is not oo
// receiver function 009,009.1

package main

import "fmt"

type deck []string

// reciever function
// print works as a function for type deck
func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}
