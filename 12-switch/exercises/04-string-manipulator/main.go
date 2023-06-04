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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {

	const (
		usage = "[command] [string]"
		l     = "lower"
		u     = "upper"
		t     = "title"
	)

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	// Create a TitleCase converter from the golang.org/x/text/cases package
	titleCase := cases.Title(language.Turkish)

	// Convert the input string to title case

	c := os.Args[1]
	switch c {
	case l:
		fmt.Printf("%s\n", strings.ToLower(os.Args[2]))
	case u:
		fmt.Printf("%s\n", strings.ToUpper(os.Args[2]))
	case t:
		fmt.Printf("%s\n", titleCase.String(os.Args[2]))
		//fmt.Printf("%s\n", strings.Title(os.Args[2]))
	default:
		fmt.Printf("Unknow command %s\n", os.Args[1])
	}
}
