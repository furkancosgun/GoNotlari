package main

import "fmt"

type User struct {
	name    string
	surname string
}

func (user *User) changeMyName() {
	user.name = "Furkan"
}

func sayMyName(user *User) {
	fmt.Println(user.name)
}

func main() {
	/*
		// defer anahtar kelimesi o an ki state sonunda çalışır bir fonksion içinde
		// kullanılıyorsa o fonksiyonun bitiminde çalıştırılır

		//Scope içerisinde once bu kod yazılmasına ragmen defer keywordu ile bu işlem ertelenir
		defer fmt.Println("Merhaba İlk Ben Yazildim")

		fmt.Println("Merhaba En Son Ben Yazildim")

		//Birden cok defer kullanımında ilk giren son cıkar mantıgı çalışmaktadır
		//Yani son defer ilk çalışcaktır
		defer fmt.Println("1")
		defer fmt.Println("2")
		defer fmt.Println("3")
		defer fmt.Println("4")
		defer fmt.Println("5")

		fmt.Println("Scope Sonuna Gelindi Artik")
	*/

	user := &User{name: "Ahmet", surname: "Kaya"}

	//Referans uzerinden erişip ismini yazdırıyoruz
	defer sayMyName(user) //Out:Furkan

	sayMyName(user) //Out:Ahmet

	//userin adını değiştirdik referans uzerinden bu durumda artık Adı Furkan
	user.changeMyName()

	//Panic fonksiyonu ile bir sistemi tamamen durdurabiliriz
	//ama bir panic ile karşılaşsak bile defer keywordu ile işaretledigimiz
	//işlemler yapılcaktır
	panic("An Error Occurred")

}
