// diagrams 4 total important concept
// passby value vs reference.
// values vs reference types

// int, bool, string, struct --- values (pointer)
// slice, maps, channels, pointers, functions -- reference

package main

import "fmt"

type contactInfo struct {
	email  string
	number int
}

type person struct {
	firstName string
	lastName  string
	zipcode   int
	contactInfo
}

func main() {
	vashishth := person{
		firstName: "Nerella",
		lastName:  "Vashishth",
		zipcode:   524201,
		contactInfo: contactInfo{
			email:  "abc@g.com",
			number: 123,
		},
	}

	//&vashishth -- gives the address of vashishth
	vashishthAddress := &vashishth

	//passing address to function
	vashishthAddress.updateFirstName("venkata")
	//it wont update as for value we need to use pointers
	vashishth.updateFirstNameWithoutPointerWontWork("ffg")
	vashishth.print()
}

//pass by value
// if we don't use pointer for value types it will create new copy and do operations on it
func (pointerToPerson *person) updateFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) updateFirstNameWithoutPointerWontWork(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
