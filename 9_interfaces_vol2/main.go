package main

import "fmt"

// Ortak bir interface oluşturma
// IShippable interfacini bir yapı implement etmek istiyorsa o yapı ShippingCost methoduna sahip olmalıdır
type IShippable interface {
	ShippingCost() int //ShippingCost adında herhangi bir parametre almayan int deger donduren bir methoda sahip
}

// Eger ki bir yapı bir interface govdesini tamamen implement ediyorsa artık
// o yapı o interface tipinden tureyebilir
type Book struct {
	desi int
}

// IShippable interfaceinin govdesini içeriyor bu durumda book yapısı aslında bir IShippable dir
func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

type Electronic struct {
	desi int
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*2
}

func main() {
	var product IShippable

	product = &Book{desi: 100}                                        //bir interface bir yapı eşlemek istiyorsak her daim referanslamamız gereklidir
	fmt.Printf("Product Shipping Cost:%v \n", product.ShippingCost()) //Out:205

	product = &Electronic{desi: 200}
	fmt.Printf("Product Shipping Cost:%v \n", product.ShippingCost()) //Out:410

	//Şekliden bir Interface listesi tanımladır bunun içerisine Electronic ve Book Yapılarımız ekleyip
	//toplu bir şekilde taşıma maliyetini hesaplayabiliriz
	products := []IShippable{&Book{desi: 10}, &Book{desi: 12}, &Book{desi: 14}, &Electronic{desi: 20}, &Electronic{desi: 25}, &Electronic{desi: 30}}
	fmt.Println(calculateTotalShippingCost(products)) //Out:267
}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}
