// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// score := 2
	fmt.Println("Enter score: ")

	// var then variable name then variable type
	var score float64

	// Taking input from user
	fmt.Scanln(&score)

	if score > 3.0 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else if score >= 2 && score <= 2.9 {
		fmt.Println("meh...")
	} else {
		fmt.Println("low")
	}
}
