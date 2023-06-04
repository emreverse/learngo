package main

import (
	"fmt"
	"time"
)

func işçi(işçi int, kuyruk chan int, cevap chan int) {
	for iş := range kuyruk {

		fmt.Println(işçi, ".işçi", iş, ".işi tamamladı")
		time.Sleep(time.Second)
		cevap <- iş

	}
}

func main() {
	işçi_sayısı := 3
	iş_sayısı := 9

	kuyruk_m := make(chan int, iş_sayısı)
	cevap_m := make(chan int, iş_sayısı)

	for i := 1; i <= işçi_sayısı; i++ {
		go işçi(i, kuyruk_m, cevap_m)
	}

	for iş_m := 1; iş_m <= iş_sayısı; iş_m++ {
		kuyruk_m <- iş_m
	}

	for a := 1; a <= iş_sayısı; a++ {
		<-cevap_m
	}
	fmt.Println("---\nTamamlandı")
}
