package main

import "fmt"

func main() {
	/*Uzunlugu belli ise array degilse slice
	//Arrayler
	//3 elemanli bir string listesi
	// var names [3]string
	// names[0] = "Pazartesi"
	// names[1] = "Sali"
	// names[2] = "Carsamba"
	// fmt.Println(names)

	//Ilk Tanımlanırkende deger verebiliriz
	var names = [3]string{"Pazartes", "Sali", "Carsamba"}
	names[1] = "Cuma" //Sonradanda index ile erişip değiştirebilirz
	fmt.Println(names)

	var dayOfWeek = [7]string{"Pazartesi", "Sali", "Carsamba", "Persembe", "Cuma", "Cumartesi", "Pazar"}
	fmt.Println(dayOfWeek[0:3]) //0.indexten başla 3.indexe kadar al 3 dahil degil
	*/
	//Slice Kullanımı herşey nerdeyse aynıdır sadece eger erişmeye calıştıgımız index boş ise sistem panic licketir
	var dayOfWeek = []string{"Pazartesi", "Sali", "Carsamba", "Persembe", "Cuma", "Cumartesi", "Pazar"}
	fmt.Println(dayOfWeek[0:3]) //0.indexten başla 3.indexe kadar al 3 dahil degil

	//Index Erişimi
	fmt.Println(dayOfWeek[6]) //Pazar
	//fmt.Println(dayOfWeek[7])   //panic:sistem durcaktır

	//Slicelara deger elemek için kullanılır
	dayOfWeek = append(dayOfWeek, "Pazardan  sonraki gun", "gibi", "gibi")
	fmt.Println(dayOfWeek)
}
