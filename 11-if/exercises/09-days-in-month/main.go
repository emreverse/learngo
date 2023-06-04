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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	leap := time.Now().Year()%4 == 0 && (time.Now().Year()%100 != 0 || time.Now().Year()%400 == 0)
	year := time.Now().Year()

	fmt.Printf("leap %v year %v\n", leap, year)

	days := 28
	fmt.Printf("Type of days %T\n", days)
	month := strings.ToLower(os.Args[1])
	//m := strings.ToLower(month)
	switch month {
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
		fmt.Printf("total days of the %v month is: %v\n", month, days)
	case "april", "june", "september", "november":
		days = 30
		fmt.Printf("total days of the %v month is: %v\n", month, days)
	case "february":
		days = 28
		fmt.Printf("total days of the %q month is: %v\n", month, days)
		if leap {
			days = 29
			fmt.Printf("total days of the %q month is: %v\n", month, days)
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
	}

}
