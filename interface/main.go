// disagrams 06 folder
// go through diagrams before going through program
// interface very tough and important concept in go core understanding

/*
   1) if we have same functionality in different functions we can make it as an interface.
   2) If receiver function implements interface then that receiver type will come under interface type
*/

package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGrettings() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// printGreeting is expecting bot type but we are sending different type
	// It's accepting because those types became part of interface by implementing definitions in interface
	printGreeting(eb)
	printGreeting(sb)
}

// here we are using same function for different types using interface
func printGreeting(b bot) {
	fmt.Println(b.getGrettings())
}

// as englishBot type implements interface, englishbot type is considered as bot type --- implicit concept
func (englishBot) getGrettings() string {
	return "hello english"
}

// as spanishBot type implements interface, spanishBot type is considered as bot type --- implicit concept
func (spanishBot) getGrettings() string {
	return "hello spanish"
}
