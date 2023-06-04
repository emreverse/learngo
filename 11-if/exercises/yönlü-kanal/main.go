package main

import "fmt"

func hoparlör(H <-chan string, ses_seviyesi int) {
	fmt.Printf("Anons ses seiyesi: %d\n", ses_seviyesi)
	fmt.Printf("Anons metni: %s\n", <-H)
}
func mikrofon(M chan<- string, mesaj string) {
	M <- mesaj
}
func main() {
	emre := make(chan string, 2)
	mikrofon(emre, "emre hoşgeldin")
	hoparlör(emre, 3)
}
