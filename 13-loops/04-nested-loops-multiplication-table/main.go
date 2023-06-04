// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

// integ := os.Args[1]

func main() {
	// integ, _ := strconv.ParseInt(os.Args[1], 10, 64)
	// integ = int(integ)
	integ, _ := strconv.Atoi(os.Args[1])
	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= integ; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= integ; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= integ; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
