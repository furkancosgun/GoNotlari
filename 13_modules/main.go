package main

import (
	"fmt"
	"myModule/helper"            //Once ana modul ismi yazılır daha sonra paket ismi yazılır
	rest1 "myModule/helper/rest" //Eger Aynı paket adından iki tane varsa bunlara farklı takma adlar verebilirz
	//rest2 "myModule/helper2/rest"
)

func main() {
	fmt.Println("main paketi")
	helper.Helper1()

	rest1.Rest1()
	//rest2.Rest2()
	helper.SayMyName() //buna erişebiliriz(buyuk harf) public
	//helper.sayMyName()//buna erişemeyiz(kucuk harf) protected
}
