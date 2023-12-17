package domain

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}
