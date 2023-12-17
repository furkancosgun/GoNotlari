package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//Ilkel veri tiplerinde go dili default olarak deger vermiştir
	// var value_int int       //default 0
	// var value_float float32 //default 0
	// var value_string string //default ""

	// fmt.Println(value_int)
	// fmt.Println(value_float)
	// fmt.Println(value_string)

	// //Referans veri tiplerinde default olarak nil(null) degeri vardır
	// var int_pointer *int //nil

	// fmt.Println(int_pointer)

	// var pointer_1 *int
	// if pointer_1 == nil {
	// 	fmt.Println("Pointer_1 Nil Herhangi Bir Yeri Gostermiyor")
	// }

	//Basit bir ornek yapmak adına bir dosya okuma işlemi yapalım os modulu ile
	//ReadFile fonksiyonu ([]byte error) olarak iki farklı deger dondurmekte
	byte_value, error_value := os.ReadFile("sample.txt")

	if error_value != nil {
		fmt.Println("Dosya Okumada Bir Hata İle Karsilasildi", error_value.Error())
	} else {
		fmt.Printf("%s", byte_value) //Eger bir hata ile karsilasmazsa dosya icindeki degeri yazdiracaktir
	}

	//Kendi fonksiyonlarımız için error dondurme
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Has Error:", err.Error())
	// } else {
	// 	fmt.Printf("Result:%v", result)
	// }

	//Egerki hata kontrolu yapmıcaksak _ tanimini verebiliriz
	result, _ := divide(10, 2)
	fmt.Println(result)

	var err error
	productServive := &ProductService{}
	product := Product{name: "", price: 100}
	err = productServive.Add(product)
	if err != nil {
		fmt.Println(err) //Product Name Is Empty
	}
	product.name = "Iphone"
	product.price = 0

	err = productServive.Add(product)
	if err != nil {
		fmt.Println(err) //Product Price Must Be Greather Than 0
	}

	product.price = 100
	err = productServive.Add(product)
	if err != nil {
		fmt.Println(err)
	}
}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

func (productService ProductService) Add(product Product) error {

	if len(product.name) == 0 {
		// return errors.New("Product Name Is Empty")//standard
		return ValidationError{code: 10, text: "Product Name Is Empty"} //Custom Error Struct Implementer error interface
	}
	if product.price <= 0 {
		// return errors.New("Product Price Must Be Greather Than 0")
		return ValidationError{code: 20, text: "Product Price Must Be Greather Than 0"}
	}
	fmt.Println("Product Successfuly Created")
	return nil
}

// Custom Error Yapısı Oluşturmak
type ValidationError struct {
	code int8
	text string
}

// Standard error interfacini implement ederiz
// standard error interfacei  Error() string methoduna sahiptir
func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Error Code:%v - Error Text:%v", validationError.code, validationError.text) //String Formatter
}

// Bir int bir de error donduren fonksiyon
func divide(s1 int, s2 int) (int, error) {
	if s2 == 0 {
		//return 0, fmt.Errorf("Zero Division Error")
		return 0, errors.New("Zero Division Error")
	}
	return s1 / s2, nil
}
