package main

import (
	"fmt"
	"time"
)

func ilk(i chan<- string) {
	time.Sleep(time.Second * 5)
	i <- fmt.Sprintf("%v", "ilk")

}

func son(s chan<- string) {

	time.Sleep(time.Second * 2)
	s <- fmt.Sprintf("%v", "son")

}

func writeTime() {
	fmt.Println(time.Now())
}

func main() {
	writeTime()
	i := make(chan string)
	s := make(chan string)
	go ilk(i)
	go son(s)

	for a := 1; a <= 2; a++ {
		select {
		case mesaj := <-i:
			fmt.Printf("%v iş tamamlandı\n", mesaj)
		case mesaj2 := <-s:
			fmt.Printf("%v iş tamamlandı\n", mesaj2)
		}
	}
	writeTime()
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printEveryTwoSeconds(ch chan<- string) {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(2 * time.Second)
// 		ch <- fmt.Sprintf("Printing every 2 seconds: %d", i)
// 	}
// 	close(ch)
// }

// func printEveryFiveSeconds(ch chan<- string) {
// 	for i := 1; i <= 3; i++ {
// 		time.Sleep(5 * time.Second)
// 		ch <- fmt.Sprintf("Printing every 5 seconds: %d", i)
// 	}
// 	close(ch)
// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go printEveryTwoSeconds(ch1)
// 	go printEveryFiveSeconds(ch2)

// 	for {
// 		select {
// 		case msg1, ok := <-ch1:
// 			if ok {
// 				fmt.Println(msg1)
// 			}
// 		case msg2, ok := <-ch2:
// 			if ok {
// 				fmt.Println(msg2)
// 			}
// 		}
// 	}
// }
