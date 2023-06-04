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

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE

	fmt.Println("there are", len(os.Args)-1, " people!")
	fmt.Println("Hello great", os.Args[1])
	fmt.Println("Hello great", os.Args[2])
	fmt.Println("Hello great", os.Args[3])
	fmt.Println("Nice to meet you all")

	str := "This is an example of a string"
	str_replacable := "This is an example of a string"
	fmt.Printf("T/F? \nDoes the string \"%s\" exist have a prefix %s", str, "Th")
	fmt.Printf("\n%t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("T/F? \nDoes the string \"%s\" exist have a prefix %s", str, "thing")
	fmt.Printf("\n%t\n", strings.HasPrefix(str, "thing"))

	fmt.Printf("T/F? \nDoes the substring \"%s\" in the main string %s", "is", str)
	fmt.Printf("\n%t\n\n", strings.Contains(str, "is"))

	fmt.Printf("T/F? \nDoes the string \"%s\" exist have a prefix %s\n", str_replacable, "example")

	if strings.Contains(str_replacable, "example") == true {
		new := strings.Replace(str_replacable, "example", "example_new", 1)
		fmt.Println(new)
	} else {
		fmt.Println("oops birşeyler yanlış gitti")
	}
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.

}
