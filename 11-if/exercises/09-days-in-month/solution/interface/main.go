package main

import "fmt"

type emre struct {
	adı  string
	türü string
}

func (e emre) isimYazdır() {
	fmt.Println(e.adı)
}

func (e emre) türYazdır() string {
	return e.türü
}

type emre_arayüz interface {
	isimYazdır()
	türYazdır() string
}

func isim(i emre_arayüz) {
	i.isimYazdır()
	fmt.Println(i.türYazdır())
}

func main() {
	t := emre{
		adı:  "emre",
		türü: "kedi",
	}
	isim(t)
	// fmt.Println()
	fmt.Printf("type of t is %T\n", t)
	// fmt.Println(t.türYazdır())

}
