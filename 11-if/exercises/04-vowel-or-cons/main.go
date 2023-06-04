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
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.args) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {

	// args := os.Args
	// ll := len(args)

	// if ll != 2 || len(args[1]) != 1 {
	// 	fmt.Println("Give me a letter")

	// } else if args[1] == "a" || args[1] == "e" || args[1] == "i" || args[1] == "o" || args[1] == "u" {
	// 	fmt.Printf("%q is a vowel.\n", args[1])
	// } else if args[1] == "w" || args[1] == "y" {
	// 	fmt.Printf("%q is sometimes a vowel, sometimes not.\n", args[1])
	// } else {
	// 	fmt.Printf("%q is a consonant.\n", args[1])
	// }

	args := os.Args
	l := len(args)
	if l != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
	} else if strings.IndexAny(strings.ToLower(args[1]), "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", args[1])
	} else if strings.ToLower(args[1]) == "w" || strings.ToLower(args[1]) == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", args[1])
	} else {
		fmt.Printf("%q is a consonant.\n", args[1])
	}
}
