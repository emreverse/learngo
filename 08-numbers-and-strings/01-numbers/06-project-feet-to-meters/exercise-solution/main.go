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

func main() {
	arg2 := os.Args[2]
	c, _ := strconv.ParseFloat(os.Args[1], 64)
	d, _ := strconv.ParseFloat(arg2, 64)
	f := c*1.8 + 32

	// Like this:
	fmt.Printf("%g ºC is %g ºF\n", c, f)
	fmt.Printf("%g ºC is %g ºF\n", d, f)

	// Or just like this (both are correct):
	fmt.Printf("%g ºF\n", f)
}
