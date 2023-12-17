package main

import (
	"fmt"
	"time"
)

func main() {
	// myChannel := make(chan int, 5)

	/*
		myChannel <- 1
		myChannel <- 2
		myChannel <- 3
		myChannel <- 4
		myChannel <- 5
		close(myChannel) //Deadlock almamak için kanal kapatılır

		//channellar uzerinde for range dongusu kurulur fakat indexe erişilemez
		//ayrıca for dongusunde surekli olarak kanal okunmaya çalışılır ta ki kanal kapanana kadar
		//bu durumda ya kanalı kapatmayıp kanala aynı hızda verilere yazacagız ya da kanalı kapatacagız
		//yoksa sistem deadlock hatası alcaktır
		for value := range myChannel {
			fmt.Println("For Range Value:", value)
		}
	*/
	/*

		//Her saniyede kanala bir veri gonderiyoruz
		go func() {
			defer close(myChannel) //Fonksiyon bittiginde donguyu kanalı kapat

			var index int
			for {
				index++
				myChannel <- index
				time.Sleep(time.Second)
				if index == 10 {
					break
				}
			}
		}()

		//Yukarıdaki dongu bittiginde kanal kapancak ve for asagıdaki for dongumuz sonlancaktır
		//Her saniye kanaldan yeni bir veri okuyoruz
		for value := range myChannel {
			fmt.Println("For Range Value:", value)
			time.Sleep(time.Second)
		}
	*/
	//Select Yapısının Kullanımı

	channel1 := make(chan string)
	channel2 := make(chan string)

	var data1 string
	var data2 string

	go func() {
		time.Sleep(10 * time.Second)
		channel1 <- "Hello"
	}()
	go func() {
		time.Sleep(time.Second)
		channel2 <- "World"
	}()

	//Boyle bir durumda channel1 e veri yazan rutin 10 saniye sonra veri yazcak
	//ama channel2 ye veri yazan channel 1 saniye sonra veri yazacak
	//bu durumda neden ben gereksiz yere channel 1 i beklemeye devam edeyim ki?
	//channel2 nin verisini alır devam ederim
	// data1 = <-channel1
	// fmt.Println("Data1:", data1)

	// data2 = <-channel2
	// fmt.Println("Data2:", data2)

	//&
	//bu gibi durumlarda select ifadesi kullanılır
	//channel1 veya channel2 hangisine ilk veri gelirse o çalışır
	//şey gibi duşunebiliirz aynı anda 5 farklı internet sitesinde hava durumu için istek attık
	//ilk hangi hava durum bilgisi bize gelirse onu kullanacagız

	//Default anahtar kelimesi ise biz select ifadesine gelene kadar en az birtane channela veri gelmesi lazım
	//eger ki gelmedi bu durumda default çalışcaktır channellardan veri gelmesi beklenmeyecektır
	select {
	case data1 = <-channel1:
		fmt.Println("Data1:", data1)
	case data2 = <-channel2:
		fmt.Println("Data2", data2)
	default:
		fmt.Println("Default")
	}

}
