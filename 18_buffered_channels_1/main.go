package main

import (
	"fmt"
)

func main() {

	/*
		//Unbufferd Channels //deadlock alır bu şekilde çalıştırılırsa
		myChannel := make(chan int)

		myChannel <- 10
	*/

	/*

		//Channellarda size verdigmiz durumlarda artık bunlar bizim için
		//Bufferd channeldir size'lar kanalın kaç veri tutabilcegini ifade eder
		myChannel := make(chan int, 2)

		myChannel <- 10
		myChannel <- 20
		myChannel <- 30 //Henuz bu kanal okunmadıgı için size'i aşacagız ve deadlock hatası alacagız

	*/
	/*
		//Boş Bir Channel Okumaya Çalışırsak Yine Deadlock Hatası Alırız
		myChannel := make(chan int, 2)
		<-myChannel
		fmt.Println("End Of The Main")
	*/

	/*Queue mantıgı


	myChannel := make(chan int, 2)
	myChannel <- 1
	myChannel <- 2

	//Ilk Giren Ilk Çıkar Mantıgında Çalışır
	fmt.Println(<-myChannel) //out:1

	myChannel <- 3           //Bir Veriyi println ile okuduk ve yazdırdık
	fmt.Println(<-myChannel) //out:2//bu seferde sıradaki gelir
	fmt.Println(<-myChannel) //out:3
	//fmt.Println(<-myChannel) //Channelda herhangi bir veri olmadıgı için :deadlock hatası alınır
	*/

	messages := make(chan string, 2)

	go func() {
		val1 := <-messages
		fmt.Println("Val1:", val1)
		val2 := <-messages
		fmt.Println("Val2:", val2)
	}()

	//Kanala Veri Gonderilir
	messages <- "Hello" //İlk Veri gonderildigi gibi go rutindeki kanal okuma block kalkar ve sonraki blocka takılır
	messages <- "World" //Ikinci blokta kalkar ve go rutin kapanır
	messages <- "Temp"  // bu degerimiz ise kanalda boş bir şekilde kalır herhangi bir çıktısını alamayız ama yukarıdakilerinin cıktısını alırız
	fmt.Println("End Of The Main")
}
