package main

import "fmt"

func sendStringsToChannel(ch chan<- string, i int, s string) {
	ch <- fmt.Sprintf("Index: %d, Value: %s", i, s)
}

func main() {
	// Create a channel to send strings
	ch := make(chan string)

	// Create a string slice
	b := []string{"emre", "yunus"}

	// Start a goroutine to send the elements of the slice to the channel
	go func() {
		for i, s := range b {
			sendStringsToChannel(ch, i, s)
		}
		close(ch)
	}()

	// Read and print the messages from the channel

	for msg := range ch {
		fmt.Println("Received message:", msg)
	}
}
