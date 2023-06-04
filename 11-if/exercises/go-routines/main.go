package main

import (
	"fmt"
	"time"
)

func harfYaz(i int) {
	harfler := "abcçd"
	for _, v := range harfler {
		fmt.Printf("%d%s", i, string(v))
	}
	fmt.Printf("\n")
}

func müzevir(dedikodu chan string) {
	dedikodu <- "emrenin diğer adı da yunus"
}

func main() {
	fmt.Println("-------Sıralı Çalışan --------------")
	for i := 0; i < 5; i++ {
		harfYaz(i)
	}

	fmt.Println("-------Eşzamanlı çalışan --------------")
	for i := 0; i < 5; i++ {
		go harfYaz(i)
	}
	time.Sleep(time.Second)

	fmt.Println("-------Kanallara geç --------------")

	kanal1 := make(chan string)
	go func() {
		kanal1 <- "selam emre"
	}()
	var1 := <-kanal1
	fmt.Println(var1)

	dedikodu := make(chan string)
	go müzevir(dedikodu)

	gazete := <-dedikodu
	fmt.Println(gazete)

	emre_kanal := make(chan int)
	go emre(emre_kanal)

	var2 := <-emre_kanal
	fmt.Printf("%v\n", var2)

}

func emre(K chan int) {
	K <- 4
}
