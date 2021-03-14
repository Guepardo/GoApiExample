package models

type Product struct {
	BasicModel
	Name  string  `json:name`
	Price float64 `json:price`
}
