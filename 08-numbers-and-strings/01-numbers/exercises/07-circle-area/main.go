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
	"math"
)

// ---------------------------------------------------------
// EXERCISE: Circle Area
//
//  Calculate the area of a circle from the given radius
//
// CIRCLE AREA FORMULA
//  area = πr²
//  https://en.wikipedia.org/wiki/Area#Circles
//
// HINT
//  For PI you can use `math.Pi`
//
// EXPECTED OUTPUT
//  radius: 10 -> area: 314.1592653589793
//
// BONUS EXERCISE!
//  1. Print the area as 314.16
//  2. To do that you need to use the correct Printf verb :)
//      Instead of `%g` verb below.
//
//    EXPECTED OUTPUT
//     radius: 10 -> area: 314.16
// ---------------------------------------------------------

func main() {
	var (
		radius = 10.
		area   float64
		pi     = math.Pi
	)

	// ADD YOUR CODE HERE
	// ...

	area = radius * pi * radius
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %g\n", radius, area)
}
