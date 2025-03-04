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
// EXERCISE: Incdecs
//
//  1. Increase the `counter` 5 times
//  2. Decrease the `factor` 2 times
//  3. Print the product of counter and factor
//
// RESTRICTION
//  Use only the incdec statements
//
// EXPECTED OUTPUT
//  -75
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THIS
	counter, factor := 45, 0.5
	// counter++
	// counter++
	// counter++
	// counter++
	// counter++
	// factor--
	// factor--
	// fmt.Println(float64(int(counter)) * factor)

	// TYPE YOUR CODE BELOW
	// ...

	for i := 0; i < 5; i++ {
		counter = counter + 1

	}
	for j := 0; j < 2; j++ {
		factor = factor - 1
	}
	fmt.Println(float64(int(counter)))
	fmt.Println(factor)
	fmt.Printf("%.f\n", float64(int(counter))*factor)

	// LASTLY: REMOVE THE CODE BELOW
	_, _ = counter, factor
}
