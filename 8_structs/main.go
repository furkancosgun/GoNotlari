package main

import "fmt"

type Addres struct {
	country string
	city    string
}

// E ticaret muşterisi
type Customer struct {
	id      int64
	name    string
	surname string
	adress  Addres
}

// Bu fonksiyon sadece customer structına baglı objeler için kullanılabilir
func (customer Customer) to_string() {
	fmt.Printf("Adi:%v , Soyadi:%v , Ulkesi:%v , Sehri:%v \n", customer.name, customer.surname, customer.adress.country, customer.adress.city)
}
func (customer Customer) changeMyName() {
	customer.name = "Ayse"
}
func (customer *Customer) changeMyNameWithRef() {
	customer.name = "Ayse" //ilkel veri tiplerinde *param_name şeklinde erişirken gelişmiş veri tiplerinde boyle bir ihityac yoktur
}
func changeCustomerName(customer Customer) {
	customer.name = "Hasan Ali"
}

func changeCustomerNameWithRef(customer *Customer) {
	customer.name = "Hasan Ali"
}

func main() {

	// var customer1 Customer
	// customer1 = Customer{id: 1, name: "Furkan", surname: "Cosgun"}

	// fmt.Println(customer1)

	// var customer2 = Customer{id: 2, name: "Ali", surname: "Ayse"}
	// fmt.Println(customer2)

	// customer3 := Customer{id: 3, name: "Ahmet", surname: "Kaya"}
	// fmt.Println(customer3)

	// customer3.name = "Mehme"
	// fmt.Println(customer3)

	//Şeklinde struct içinde struct tanımlaması yapabilirz
	customer := Customer{id: 1, name: "Furkan", surname: "COSGUN", adress: Addres{country: "Turkiye", city: "Nevsehir"}}
	fmt.Println(customer)

	//Şeklinde Ozellestirlimiş fonksiyonu kullanabilir
	customer.to_string()
	changeCustomerName(customer) //Strutlarda pass by value mantıgında çalışmaktadır
	customer.to_string()
	changeCustomerNameWithRef(&customer) //ancak pointer kullanarak degeri degisitrebiliriz
	customer.to_string()
	customer.changeMyName() //Referans kullanamadan structin sahip oldugu fonksyon ile de degişim saglayamayız
	customer.to_string()
	customer.changeMyNameWithRef() //referans kullanarak sturctın sahip oldugu fonk ile degişim saglayabilirz
	customer.to_string()
}
