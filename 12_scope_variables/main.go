package main

import "fmt"

// Global scope
const name = "furkan"

func main() {

	// if true {
	// 	var x = 100 //block scope
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// var x = 100 //function scope
	// if true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// sayMyName() //function scope
	// fmt.Println(name)

	fmt.Println(name) //Global scope tan erişildi//Out:furkan

	var name = "ali"  //hem globalde hemde localde değişken varsa aynı isimde once localdekini baz alır
	fmt.Println(name) //Out:ali

	fmt.Println(name)  //Out:ali
	setMyName("hasan") //Out:hasan
	fmt.Println(name)  //Out:ali
}

func setMyName(newName string) {
	var name = newName //Localdekini değiştiri sadece
	fmt.Println(name)
}

func sayMyName() {
	//fucntion scope burası burdaki degisken bu fonksiyon bittikten sonra bellekten silinir
	var name = "furkan"
	fmt.Println(name)
}
