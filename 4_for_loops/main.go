package main

import "fmt"

func main() {

	//Şeklinde de for dongusunde direk olarak tanımlayabiliriz 'i' degerini
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	index := 0
	for index <= 100 {
		fmt.Printf("Index: %v \n", index)
		index++
	}

	//Sonsuz dongu
	for {
		fmt.Println("Sonsuz Dongu")
		break //Donduden Çıkartır
		//continue //dongu içerisinde aşagıdaki koda devam etmez dongu başına gonderir
	}

	names := []string{"Furkan", "Hasan", "Ali", "Ahmet"}
	for index, name := range names {
		fmt.Printf("Index:%v , Value:%v \n", index, name)
		fmt.Printf("Index:%v , Value:%v \n", index, names[index]) //Şeklinde de kullanabilir bir loop açıp
	}

	second_index := 0
	for second_index < len(names) { //Liste uzunlugu kadar bir dongu kurulur
		fmt.Printf("Index:%v,Value:%v \n", second_index, names[second_index])
		second_index++
	}
}
