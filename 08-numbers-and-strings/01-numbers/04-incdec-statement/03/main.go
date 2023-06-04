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
	"reflect"
	"strconv"
)

// incdec is a statement

func main() {
	var counter int

	// following "statements" are correct:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	counter-- //
	counter++ // 4
	counter++
	fmt.Printf("There are %d line(s) in the file\n",
		counter)
	//q := strconv.Itoa(counter)
	fmt.Println("There are", strconv.Quote(strconv.Itoa(counter)), "lines(s) in the file")
	// the following "expressions" are incorrect:

	// counter = 5+counter--
	// counter = ++counter + counter--
	emre := 0

	for i := 0; i < counter; i++ {
		emre++
		fmt.Println("The value of count emre is:", strconv.Quote(strconv.Itoa(emre)))
		// 	time.Sleep(time.Second)
	}

	nokta := [3][2]int{{11, 12}, {21, 22}, {31, 32}}
	fmt.Println(reflect.TypeOf(nokta))
	fmt.Println("Dizinin tamamı", nokta)
	fmt.Println("3x2 matris biçiminde:")
	for satır := 0; satır < 3; satır++ {
		for sütun := 0; sütun < 2; sütun++ {
			fmt.Print(nokta[satır][sütun], " ")
		}
		fmt.Println()
	}
}
