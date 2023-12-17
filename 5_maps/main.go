package main

import "fmt"

func main() {
	//Maps:Key value şeklinde deger tutan listeler diyebilir

	//Key degerim string value degerim int
	// var sozluk map[string]int

	// sozluk = make(map[string]int, 0)
	// sozluk["one"] = 1
	// sozluk["two"] = 2
	// sozluk["three"] = 3
	// fmt.Println(sozluk)           //Butun bir mapi gosterir
	// fmt.Println(sozluk["one"])    //buldugu degeri doner
	// fmt.Println(sozluk["onetwo"]) //bulunmayan bir degerde 0 olarak verir

	//Şeklindede kısaca tanımlayabiliriz
	sozluk := make(map[string]int, 0)
	sozluk["one"] = 1
	sozluk["two"] = 2
	sozluk["three"] = 3
	fmt.Println(sozluk)

	//Şeklindede kısaca deger vererek tanımlayabilirz
	sozluk2 := map[string]int{"one": 1, "two": 2, "three": 3}
	delete(sozluk2, "one") //Sozluk2 içerisinden one key alanlı degeri siler
	fmt.Println(sozluk2)
}
