package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("başlıyorum, 3 saniye sürecek")
	timer := time.NewTimer(3 * time.Second)
	go func() {
		fmt.Println("başladımmmm")
		<-timer.C
		fmt.Println("tamamlandı")
	}()
	time.Sleep(time.Second * 4)
	cancel := timer.Stop()
	fmt.Println("Durduruldu mu?:", cancel)

	// periyodicly doing a work with Ticker

	ticker := time.NewTicker(time.Second * 4)
	go func() {
		fmt.Println("her 4 saniye bir iş yapabilirim")
		for time := range ticker.C {
			fmt.Printf("Yapılan iş %s\n", time)
		}
	}()
	time.Sleep(time.Second * 16) // make it 17 to print 4 lines of each 4-second job
	ticker.Stop()

}
