package main

import "fmt"

func main() {
	var age = 17
	if age >= 18 {
		fmt.Println("Yasiniz 18'den buyuktur ")
	} else {
		fmt.Println("Yasiniz 18'den kucuktur")
	}

	a := 1
	b := 2
	c := 3
	if a < b && b < c && c < 4 {
		fmt.Println("A kucuktur B den ,B kucuktur C den,C kucuktur 4 ten")
	}

	if a < b || b > 99 {
		fmt.Println("A kucuktur B den veya B buyuktur 99")
	}
}
