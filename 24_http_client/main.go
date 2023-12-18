package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Currency struct, TCMB API'den alinacak XML verisinin yapisini temsil eder.
type Currency struct {
	CurrencyCode string `xml:"CurrencyCode,attr"`
	CurrencyName string `xml:"CurrencyName"`
	ForexBuying  string `xml:"ForexBuying"`
	ForexSelling string `xml:"ForexSelling"`
}

// TCMBResponse struct, TCMB API'den alinacak XML verisinin genel yapisini temsil eder.
type TCMBResponse struct {
	Currencies []Currency `xml:"Currency"`
}

func main() {
	// TCMB API URL
	apiURL := "https://www.tcmb.gov.tr/kurlar/today.xml" // TCMB'nin günlük XML API'si

	// HTTP GET isteği gönderme
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("HTTP GET isteği başarisiz:", err)
		return
	}
	defer response.Body.Close()

	// HTTP yanitini okuma
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("HTTP yaniti okunamadi:", err)
		return
	}

	// XML verisini TCMBResponse struct'ina çözme
	var tcmbResponse TCMBResponse
	err = xml.Unmarshal(body, &tcmbResponse)
	if err != nil {
		fmt.Println("XML çözme hatasi:", err)
		return
	}

	// Kurlari ekrana yazdirma
	fmt.Println("Günlük Döviz Kurlari:")
	for _, currency := range tcmbResponse.Currencies {
		fmt.Println(currency)
	}
}
