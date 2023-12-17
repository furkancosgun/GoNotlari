package main

import (
	"context"
	"fmt"
	"time"
)

var productChannel = make(chan Product)

func main() {

	ctx, cancelContext := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelContext()

	go GetProductDetailFromExternalServer(100)

	select {
	case productDetails := <-productChannel:
		fmt.Println("External Servisten Cevap Geldi")
		fmt.Println(productDetails)
	case <-ctx.Done():
		fmt.Println("ERR:Time Out")
	}

	fmt.Println("End Of The Main")

}

type Product struct {
	id   int
	name string
}

func GetProductDetailFromExternalServer(id int) {
	time.Sleep(5 * time.Second)
	productChannel <- Product{id: id, name: "Tv"}
}
