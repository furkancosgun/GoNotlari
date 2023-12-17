package main

import "fmt"

func main() {

	/*
		//Sıradan bir değişken
		var a int
		a = 100

		//Adress tutan bir değişken(pointer)
		var p *int
		p = &a

		fmt.Println("A'nin degeri:", a)
		fmt.Println("A'nin adresi:", p)            //A nin adresini verir
		fmt.Println("A'nin adresteki degeri:", *p) //A'nin adresine gider ve veriyi getirir
		*p = 200                                   //A'nin adrsine gider ve degeri degiştirir
		fmt.Println("A'nin degeri:", a)
	*/

	var a = 10
	var b int
	var p *int

	b = a   //A'nin degerini B ye atar//bu kısımda sadece deger aktarımı saglanır adres veya referans aktarımı saglanmaz
	p = &b  //B nin adresini P pointerina atar
	*p = 20 //B nin adresindeki degeri değiştirir
	fmt.Printf("A'nin degeri:%v,B'nin degeri:%v \n", a, b)

	var c int = 10
	fmt.Println(c) //10
	add10(c)       //burdaki işlem aslıdna c yi gondermek yerine c degerini gondermektir o yuzden cıktımız degismez
	fmt.Println(c) //10

	add20(&c)      //C nin adresini gonderdim adresinden verisini degistirecegiz
	fmt.Println(c) //30

	//go dilindeki int bool string gibi degiskenler ilkel veri tipleridir
	//slice ve arrayler referans veri tipleridir bir fonksyonda pass by value yerine pass by referance mantıgında gitmektedir
	var numbers = []int{1, 2, 3}
	fmt.Println(numbers) //1,2,3
	changeValue(numbers) //pointer kullanmadan deger degiscektir
	fmt.Println(numbers) //100,2,3
}
func changeValue(numbers []int) {
	numbers[0] = 100
}
func add20(x *int) { //pass by referance
	*x = *x + 20
}

func add10(x int) { //pass by value
	x += 10
}
