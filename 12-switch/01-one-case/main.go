// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	city := "Paris"
	emre := city
	fmt.Println("şehir adı giriniz")
	fmt.Scanf("%s", &city)
	fmt.Println("sonuç ekrana yazdırılıyor")
	switch city {
	case emre:
		fmt.Println("France")
	case city:
		fmt.Printf("%s\n", city)
		// default:
		// 	fmt.Printf("%s", "default değer")
	}

	// SIMILAR TO IF
	// ------------------------------------

	// switch statement is converted to an if statement
	// automatically behind the scenes
	//
	// above switch statement is similar to this if

	// if city == "Paris" {
	// 	fmt.Println("France")
	// }
}
