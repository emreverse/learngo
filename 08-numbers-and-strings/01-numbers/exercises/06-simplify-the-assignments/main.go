// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Simplify the Assignments
//
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	width += 1
	width += height
	width -= 1
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
	fmt.Print(width, "\n")
}
