package main

import "fmt"

// Basit bir encoder arayuzu
type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct{}

func (xEncoder XEncoder) Encode(value string) {
	fmt.Println("XEncoder Ile Encode Edildi")
}

func (xEncoder XEncoder) Decode(value string) {
	fmt.Println("XEncoder Ile Decode Edildi")
}

type YEncoder struct{}

func (yEncoder YEncoder) Encode(value string) {
	fmt.Println("YEncoder Ile Encode Edildi")
}
func (yEncoder YEncoder) Decode(value string) {
	fmt.Println("YEncoder Ile Decode Edildi")
}

func main() {

	//Once Bu Yapıya Karar verdik
	// xEncoder := &XEncoder{}
	// xEncoder.Encode("123") //Şifreleme yap
	// xEncoder.Decode("q1e") //Şifrelemeyi çoz

	//Sonra Buna geçtik ama butun yazılımda encoder ve decoder kullandıgımız her yeri degiştirmemiz gerekti
	// yEncoder := &YEncoder{}
	// yEncoder.Encode("123") //Şifreleme yap
	// yEncoder.Decode("q1e") //Şifrelemeyi çoz

	///bunun yerine alttaki gibi interface yapısı kullanıldı sadece encoder referansının yaratıldıgı yerde degiştirildi
	var encoder IEncoder
	// encoder = &XEncoder{} //Gibi
	encoder = &YEncoder{}
	encoder.Encode("123")
}
