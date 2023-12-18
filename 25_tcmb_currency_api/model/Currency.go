package model

type Currency struct {
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
}
type TCMBCurrencyList struct {
	Currencies []Currency `xml:"Currency"`
}
