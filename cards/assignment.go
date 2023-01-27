package main

import "fmt"

func evenOrOdd(arr []int) {
	for _, number := range arr {
		if number%2 == 0 {
			fmt.Println(fmt.Sprint(number) + " " + "is even")
		} else {
			fmt.Println(fmt.Sprint(number) + " " + "is odd")
		}

	}
}
