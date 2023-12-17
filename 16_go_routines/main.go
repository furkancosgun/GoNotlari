package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		fmt.Println("Start Of The Main")

		//Main Rutin Uzerinde Çalışır
		// sayMyName("Furkan")

		//Yeni bir go rutini açar ve orda çalışır
		//yapılan işlem ne kadar uzun olursa olsun beklenmez ve sonra kodlara devam eder
		// go sayMyName("Cosgun") //Çıktısı yok

		//time.Sleep(time.Second * 3) //sistemi 3 saniye bekletiriz yukarıda çağırdıgımız go rutini bitebilir ve çıktısını gorebiliriz

		wg := sync.WaitGroup{} //go rutinleri için bekleme grubu
		wg.Add(2)              //Kac rutin bekleyecegimizi soyleriz

		go func() { //Anonymous bir go rutin fonksyonu yaratılır
			sayMyName("1.Rutinim Ben") //içine yapcagımız işlem eklenir
			wg.Done()                  //yapcagımız işlem bitince waigroup a haber verilir ben bittim diye(1 go rutini bitmiş olur)
		}()

		//aynı mantıkla ikincki fonksiyon bu da haberir verir ben bittim diye
		go func() {
			sayMyName("2.Rutinim Ben")
			//wg.Done()
			//wg.Add(-1) //Done veya add(-1) aynı mantıga gelmektedir
			defer wg.Done() //ile de kullanabiliriz ki bu durumda da fonksyonun sonuna gelince cagrılacaktır
		}()

		fmt.Println("Go Rutinler Aktif Edildi")

		//Rutinler bitene kadar da kodun aşagı satırlara inmesi beklenir aslında mantık bu kadar basittir
		//rutin içinde bir sayaç mantıgında bekletilir
		wg.Wait()

		//egerki bekleyecegimiz rutin sayısı done etcegimiz rutin sayısından fazla ise
		//o zaman deadlock hatası alıncaktır bu da sonsuza kadar beklemek anlamına gelmektedir
		//bu gibi durumlarda go dili bizi uyarır

		//egerki bekleyecegimiz rutin sayısı done etcegimiz rutin sayısından az ise
		//bu durumda waitGroup negative counter hatası alcagız bu da uygulamamıza panic vercektir

		//egerki main rutin biterse butun rutinler biter ama egerki

		fmt.Println("End Of The Main")
	*/

	// startTime := time.Now() //Başlangıc Zamanı Alınır

	// func() {
	// 	fmt.Println("Function 1")
	// 	time.Sleep(time.Second * 3) //3 saniye bekler
	// }()
	// func() {
	// 	fmt.Println("Function 2")
	// 	time.Sleep(time.Second * 3) //3 saniye bekler
	// }()
	// func() {
	// 	fmt.Println("Function 3")
	// 	time.Sleep(time.Second * 3) //3 saniye bekler
	// }()

	// //Başlangıctan beri geçen süreyi yazar(9 saniye)
	// fmt.Println("Passed Time:", time.Now().Sub(startTime))

	startTime := time.Now()

	//Bekleme Grubu Oluştrulur
	wg := sync.WaitGroup{}

	//3 go rutini bekleyecegimiz soylenir
	wg.Add(3)

	go func() {
		fmt.Println("Function 1")
		time.Sleep(time.Second * 3) //3 saniye bekler
		wg.Done()                   //go rutinin  bittigi soylenir
	}()
	go func() {
		fmt.Println("Function 2")
		time.Sleep(time.Second * 3) //3 saniye bekler
		wg.Done()                   //go rutinin  bittigi soylenir
	}()
	go func() {
		fmt.Println("Function 3")
		time.Sleep(time.Second * 3) //3 saniye bekler
		wg.Done()                   //go rutinin  bittigi soylenir
	}()

	wg.Wait() //go rutinlerinin bitmesi beklenir

	//Başlangıctan beri geçen süreyi yazar(3 saniye)
	fmt.Println("Passed Time:", time.Now().Sub(startTime))

}

func sayMyName(name string) {
	fmt.Println(name)
}
