package main

import "fmt"

func main() {

	//Channelar go rutinler arasındaki iletişimi saglar

	//go rutinlerde async olarak çalıştıkları için donuş ifadeleri kullanmaz

	//Şeklinde bir donuş ifadesi alamayız
	// value := go fn1()

	/*
		//make diyerek bir string ifade alabilen channel oluşturuyoruz
		myChannel := make(chan string)

		go func() { //Kanala veri gondermek
			//Egerki kanala veri gondermezsek deadlock hatası alırız
			myChannel <- "Kucuktur Tire(<-) işareti ile bu kanala veri gonderebiliriz"
		}()

		//Kanaldan veriyi almak :(unbuffered kanallarda kanala veri gelene kadar kod burda bekletilir)
		//Kanal açık mı kapalı mı
		channelMessage, isOpen := <-myChannel

		fmt.Println(channelMessage, isOpen)
	*/
	/*
		//Yeni bir kanal oluştrulur
		myChannel := make(chan string)

		//Eger ki kanala veri gonderirken hiçbir rutin bu kanalı dinlemiyorsa
		//o zaman yine deadlock hatası alırız
		//bir kanal yapısında hem yazanın hemde okuyanın olması gerekir
		myChannel <- "Kanala Veri Gonderiyorum"
	*/

	//Channel Oluşturlur
	myChannel := make(chan string)
	//Go rutinleri beklemek için illaki waitGroup oluşturmaya gerek yok
	//bunu channeler vasıtası ile de yapabiliriz
	doneChannel := make(chan bool)

	go func() {
		//Channeldan Mesaj Okunuyor
		message := <-myChannel
		fmt.Println(message)
		//Main rutin tarafından beklenen channela veri gonderiliyor
		doneChannel <- true
	}()

	go func() {
		//Channela veri gonderiliyor
		myChannel <- "Mesaj Gonderiliyor"
	}()

	//doneChanneldan veri gelmesi bekleniyor
	//burdan gelen veriyi illaki bir degiskene atmak veya kullanmak zorunda degiliz
	//o zaman bu durumda aşagıdaki yapıyı kullanabailirz
	<-doneChannel
	fmt.Println("End Of The Main")

}

func fn1() int {
	return 1
}
