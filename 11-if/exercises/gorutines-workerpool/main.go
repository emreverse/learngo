package main

import (
	"fmt"
	"time"
)

func işçi(işçi_no int, iş_kuyruğu chan int, iş_cevapları chan int) {
	for iş := range iş_kuyruğu {
		fmt.Println(işçi_no, " .işçi", iş, ". işi yapıyor")
		time.Sleep(time.Second * 1)
		iş_cevapları <- iş
	}
}

func main() {

	işçi_sayısı := 3
	iş_sayısı := 9
	iş_kuyruğu := make(chan int, iş_sayısı)
	iş_cevapları := make(chan int, iş_sayısı)

	for i := 1; i <= işçi_sayısı; i++ {
		go işçi(i, iş_kuyruğu, iş_cevapları)
	}

	for iş := 1; iş <= iş_sayısı; iş++ {
		iş_kuyruğu <- iş

	}

	for a := 1; a <= iş_sayısı; a++ {
		<-iş_cevapları
	}
	fmt.Println("----\nTüm işler tamamlandı.")
}
