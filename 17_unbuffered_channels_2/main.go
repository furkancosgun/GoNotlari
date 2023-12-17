package main

import "fmt"

func main() {
	/*
		//Bir Unbuffered kanal oluştrulur
		myChannel := make(chan int)

		go func() {
			//1 den 10 a kadar dongu kurulur ve her 1 saniyede bir kanala dongudeki i degeri yollanır
			for i := 0; i <= 10; i++ {
				myChannel <- i
				fmt.Printf("Veri Gonderildi:%v \n", i)
				time.Sleep(time.Second)
			}
			//dongu bittikten sonra
			//Kanalı Kapatırız
			close(myChannel)
		}()

		//Main rutinde bir sonsuz dongu yaratırız
		for {
			//Kanaldan gelen veriyi surekli olarak okur ve kanalın hala açık olup olmadıgını kontrol ederiz
			//eger ki kanal kapalıysa biz de donguyu durdururz
			value, isOpen := <-myChannel
			if !isOpen {
				break
			}
			fmt.Printf("Alinan Veri:%v \n", value)
		}
	*/

	/*
		//Her saniyede bir ekrana gunun saatini yazdırır
		myChannel := make(chan time.Time)

		go func() {
			for {
				myChannel <- time.Now()
				time.Sleep(time.Second)
			}
		}()

		for {
			value := <-myChannel
			fmt.Println(value.Clock())
		}
	*/

	myChannel := make(chan int)

	go func() {
		value := <-myChannel
		fmt.Println("Value Took By Func1:", value)
	}()
	go func() {
		value := <-myChannel
		fmt.Println("Value Took By Func2:", value)
	}()
	go func() {
		value := <-myChannel
		fmt.Println("Value Took By Func3:", value)
	}()

	//Kanala veri gonderdikten sonra rutinlerden biri çalışcaktır
	myChannel <- 10
	fmt.Println("End Of The Main")
}
