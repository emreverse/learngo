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
	"strings"
)

func main() {
	city := os.Args[1]
	city = strings.ToLower(city)
	switch city {
	case "paris":
		fmt.Println("France")
		// break // unnecessary in Go

		// vip := true
		// fmt.Println("VIP trip?", vip)

	case "tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Default")
		// fmt.Println("VIP trip?", vip)
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
