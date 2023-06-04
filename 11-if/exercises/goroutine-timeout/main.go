package main

import (
	"fmt"
	"time"
)

func süre_tut(s chan bool, a time.Duration) {
	time.Sleep(time.Second * a)

	s <- true

}

func iş_yap(i chan bool, a time.Duration) {
	time.Sleep(time.Second * a)
	i <- true

}

func main() {
	var iş time.Duration
	var süre time.Duration
	fmt.Println("Tamamlanılması gereken iş süresini girin")
	fmt.Scan(&iş)
	fmt.Println("Beklenilmesi gereken süreyi giriniz")
	fmt.Scan(&süre)

	s := make(chan bool, süre)
	i := make(chan bool, iş)
	go iş_yap(i, iş)
	go süre_tut(s, süre)
	fmt.Print("Sonuç ekrana yazdırılıyor\n")
	select {
	case <-i:
		fmt.Printf("iş beklenen %d saniyede yapıldı\n", iş)
	case <-s:
		fmt.Printf("zaman aşımı oldu, iş beklenen %d saniyede tamamlanamadı", süre)
	}

}
