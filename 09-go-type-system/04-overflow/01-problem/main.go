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
	var (
		width  uint8 = 254
		height       = 255 // int
	)
	fmt.Println("width:", width)
	width++
	fmt.Println("width:", width)
	fmt.Println("int width:", int(width))
	fmt.Println("uint8 height:", uint8(height))
	if width < uint8(height) {
		fmt.Println("height is greater")
	} else if width >= uint8(height) {
		fmt.Println("width is greater")
	}
	width++
	fmt.Println("width:", width)
	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height)
}
