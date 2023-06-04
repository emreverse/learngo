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
	city := strings.ToLower(os.Args[1])

	switch {
	case city == "istanbul", city == "ardahan":
		fmt.Printf("%s amazingly  ", city)
		fallthrough
	case city != "istanbul", city != "ardahan":
		fmt.Print("beautiful ")
		fallthrough
	default:
		fmt.Println("bir şehir")

	}

}
