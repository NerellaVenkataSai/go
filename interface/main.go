// disagrams 06 folder
// go through diagrams before going through program
// interface very tough and important concept in go core understanding

/*
   1) if we have same functionality in different functions we can make it as an interface.
   2) If receiver function implements interface then that receiver type will come under interface type
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGrettings() string
}

type customWriter struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// printGreeting is expecting bot type but we are sending different type
	// It's accepting because those types became part of interface by implementing definitions in interface
	printGreeting(eb)
	printGreeting(sb)

	//http read oand write interface
	fetchApi()
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

func fetchApi() {
	cw := customWriter{}
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("error while fetching")
		os.Exit(1)
	}

	// using read interface in resp
	/*
		Read interface is expecting byte slice as param as it reference type it's updating it
		so we are using it in next line 66 to print the repsonse body
	*/
	bod := make([]byte, 9999)
	resp.Body.Read(bod)
	// fmt.Println(string(bod))

	// standard way to print res
	fmt.Println("------------------------------------------")
	// io.Copy(os.Stdout, resp.Body)

	// custom writer interface using in copy
	io.Copy(cw, resp.Body)
}

// Write reciever function is implementing same Write interface
// so customWriter struct became part of Write interface

func (customWriter) Write(bs []byte) (n int, err error) {

	fmt.Println("custome write", string(bs))

	return len(bs), nil
}
