package main

import "fmt"

type Book struct {
	desi int
}

// bookun shipping cost hesaplamsı farkı
func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

type Electronic struct {
	desi int
}

// electronic yapısının hesaplamsı farklı
func (electronic Electronic) ShippingCost() int {
	return 10 + electronic.desi*2
}

func main() {
	// book1 := Book{desi: 50}
	// book2 := &Book{desi: 25}

	// fmt.Println(book1.ShippingCost())
	// fmt.Println(book2.ShippingCost())

	//Egerki bir slice veya array tanımlıyorsa ve tip olarka bir struct verdiysek buna deger verirken
	//tekrardan Book{desi:..} gibi yazmamıza gerek yoktur
	// bookList := []Book{{desi: 10}, {desi: 20}, {desi: 30}, {desi: 40}, {desi: 50}}
	// fmt.Println(calculateTotalShippingCost(bookList)) //Out:150

	electronic1 := Electronic{desi: 100}
	fmt.Println(electronic1.ShippingCost())
}

// Eger interfaceleri kullanmazsak benzer yapılarda surekli fonksyon yazmamız gerekir
func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost() //Her iki yapıda ShippingCost methodunda sahip bu durumda bunları tek bir interface uzerinden yonetebilriz
	}
	return total
}
func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.ShippingCost()
	}
	return total
}
