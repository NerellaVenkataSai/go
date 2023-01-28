// diangrams 05 folder
// check difference bwtween structs and maps
package main

import "fmt"

type colors map[string]string

func main() {

	colorsMap := colors{
		"black": "#000",
		"white": "#fff",
	}

	colorsMap["red"] = "#rrr"
	delete(colorsMap, "black")

	colorsMap.printColorValues()

}

func (c colors) printColorValues() {
	for _, value := range c {
		fmt.Println(value)
	}
}
