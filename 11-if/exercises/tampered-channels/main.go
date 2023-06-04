package main

import (
	"fmt"
)

func main() {

	emre_tampered := make(chan string, 5)
	tamponlu_kanal := make(chan string, 3)
	tamponlu_kanal <- "tamponlu"
	tamponlu_kanal <- "kanal"

	// fmt.Println(<-tamponlu_kanal)
	// fmt.Println(<-tamponlu_kanal)
	close(tamponlu_kanal)
	for i := range tamponlu_kanal {
		fmt.Println(i, "1")
	}

	b := []string{"emre", "yunus"}
	fmt.Println("length of slice is:", len(b))
	go func() {
		for _, s := range b {
			emre_tampered <- s

		}
		close(emre_tampered)
	}()
	for msg := range emre_tampered {
		fmt.Println(msg)
	}

	// Using with different function
	emre_tampered2 := make(chan string, 7)
	go sendToChannel(b, emre_tampered2)
	for msg2 := range emre_tampered2 {
		fmt.Println(msg2)
	}

}

func sendToChannel(a []string, ch chan<- string) {
	for _, value := range a {
		ch <- value
	}
	close(ch)
}
