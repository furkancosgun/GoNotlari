package helper

import "fmt"

const GLOBAL_NAME = "FURKAN"

// ıkı fonk tanımladık birisi buyuk harf ile başlıyor digeri kucuk harf ile
// buyuk harfle başlayanlar public(heryerden erişilebilir)
// kuçuk harf protected(aynı paket içinde kullanılabilir sadece)
func SayMyName() {
	fmt.Println(GLOBAL_NAME)
}

func sayMyName() {
	fmt.Println("FURKAN COSGUN")
}
