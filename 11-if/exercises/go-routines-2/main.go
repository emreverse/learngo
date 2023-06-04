package main

import (
	"fmt"
	"time"
)

func bekle() {
	fmt.Scanln()
	fmt.Println("Btiyor")

}

func sensordenKanala(K chan<- string) {
	for {

		K <- "zaman: " + time.Now().String()

	}
}

func kanaldanEkrana(K <-chan string) {

	for msg := range K {
		fmt.Println(msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {

	m := make(chan string)
	fmt.Println("Çıkmak için Enter tuşuna basınız")
	go kanaldanEkrana(m)
	go sensordenKanala(m)

	bekle()
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func bekle() {

// 	var enterTuşu string
// 	fmt.Scanln(&enterTuşu)
// 	fmt.Println("Btiyor")

// }

// func main() {
// 	// Create a channel to send and receive strings
// 	ch := make(chan string)

// 	// Start a goroutine to send the current time as a string through the channel every second
// 	go func() {
// 		for {
// 			ch <- time.Now().String()

// 		}
// 	}()

// 	// Start a goroutine to read the messages from the channel
// 	go func() {
// 		for msg := range ch {
// 			fmt.Println("Received:", msg)
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()

// 	// Wait for some time to allow the goroutines to run
// 	bekle()
// }
