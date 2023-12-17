package main

import "fmt"

func main() {

	standardFunction("QWE")

	func(text string) { //Anonymous function
		//Yukarıdaki standard fonksiyonla aynı işleve sahiptir
		fmt.Println(text)
	}("QWE")

	//Bir değişkene de fonksyon gorevini atayabiliriz
	//Anonymous function
	anonymousFunction := func(text string) {
		fmt.Println(text)
	}
	anonymousFunction("ASD")

	//Bir hesap makinesi mantıgında kullanabiliriz
	//iki int parametre ve (bir 2 int parametre alan int donuş degeri olan) bi fonksiyon alır
	calculator(100, 100, sum)

}

func sum(s1 int, s2 int) int {
	return s1 + s2
}

// iki int parametreleri ve int donuş tipine sahip bir fonksiyon alır
func calculator(s1 int, s2 int, fn func(int, int) int) {
	fmt.Println("Calculator Started")
	defer fmt.Println("Calculator Finished") //Scope sonunda çalışacak

	fmt.Println("Result:", fn(s1, s2))
}

func standardFunction(text string) {
	fmt.Println(text)
}
