package helper

import "fmt"

func Helper1() {
	fmt.Println("Helper1 Paketi")

	//Package scope uzerinden erişti eger aynı paket içidne birden cok go dosyası varsa
	//bunlar birbir arasında hiçbir import kullanmadan erişilebilir
	fmt.Println(GLOBAL_NAME)

	sayMyName()
}
