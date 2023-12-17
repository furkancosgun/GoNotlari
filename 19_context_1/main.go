package main

import (
	"context"
	"fmt"
	"time"
)

var productChannel = make(chan Product)

func main() {

	//Context:Context bir eşzamanlı işlerin yönetiminde kullanılan built-in bir paket.
	//Biz bunu bir servis istegi atmak uzerinde kullancagiz servise istek atıp belirttigimiz
	//surede gelmezse o istegi beklemeyecegiz ve uygulamamızı sonlandıracagız

	//context.WithTimeout diyere bir zaman kapsullu contextimizi elde ediyoruz
	//Ilke paremetre parent context :context.Background() diyerek boş bir context nesnesi yaratabiliriz
	//2. parametre ise o contextten donen channelin kaç saniye sonra biz yanıt vercegidir

	//donuşlerden ilk parametre contextin kendisi ,ikinci parametre o contexti iptal etmek için kullandıgımız
	//fonksiyon
	ctx, cancelContext := context.WithTimeout(context.Background(), 3*time.Second)
	//Main fonksiyonu bitince context ile işim kalmıyor o da kapanabilir
	defer cancelContext()

	//Servise istegimizi attık 5 saniye sonra donecek
	go GetProductDetailFormExternalService(10)

	//Servisten Cevabın Gelmesi Beklenir ve context nesnesi ile timeout kontrolu yaparız
	select {
	case productDetail := <-productChannel: //Servis istegi bekleme
		fmt.Println("Urun Detaylari Geldi")
		fmt.Printf("Product Id:%v - Product Name:%v", productDetail.id, productDetail.name)
	case <-ctx.Done(): //context yonetimindeki time out kanalı bu kanal bize cevabı yukarıda ilk oluştudugumuz sure sonunda donecektir
		fmt.Println("time out")
	}

	fmt.Println("End Of The Main")
}

type Product struct {
	id   int
	name string
}

// Dış Servis Ornegi Olsun Diye Bu Fonksiyonu 5 Saniye Beklettik
func GetProductDetailFormExternalService(id int) {
	time.Sleep(time.Second * 5)
	productChannel <- Product{id: id, name: "Iphone 12"}
}
