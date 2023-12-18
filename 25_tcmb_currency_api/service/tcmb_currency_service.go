package service

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"
	"tcmb_currency_api/model"
)

type ICurrencyService interface {
	GetAllCurrencies() ([]model.Currency, error)
	GetCurrencyByCurrCode(currcode string) (model.Currency, error)
}

type TCMBCurrencyService struct{}

func New() ICurrencyService {
	return &TCMBCurrencyService{}
}

func (currencyService *TCMBCurrencyService) GetAllCurrencies() ([]model.Currency, error) {
	var tcmbResponse model.TCMBCurrencyList

	// TCMB API URL
	apiURL := "https://www.tcmb.gov.tr/kurlar/today.xml" // TCMB'nin günlük XML API'si

	// HTTP GET isteği gönderme
	response, err := http.Get(apiURL)
	if err != nil {
		log.Printf("HTTP GET isteği başarisiz: %v\n", err)
		return tcmbResponse.Currencies, err
	}
	defer response.Body.Close()

	// HTTP yanitini okuma
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("HTTP yaniti okunamadi: %v\n", err)
		return tcmbResponse.Currencies, err
	}

	// XML verisini TCMBResponse struct'ina çözme
	err = xml.Unmarshal(body, &tcmbResponse)
	if err != nil {
		log.Printf("XML çözme hatasi: %v\n", err)
		return tcmbResponse.Currencies, err
	}

	return tcmbResponse.Currencies, nil
}

func (currencyService *TCMBCurrencyService) GetCurrencyByCurrCode(currcode string) (model.Currency, error) {
	var currency model.Currency

	currencies, err := currencyService.GetAllCurrencies()
	if err != nil {
		return currency, err
	}
	for _, curr := range currencies {
		if curr.CurrencyCode == currcode {
			currency = curr
			break
		}
	}

	// Eğer boşsa
	if (currency == model.Currency{}) {
		return currency, errors.New("Para birimi bulunamadi")
	}
	return currency, nil
}
