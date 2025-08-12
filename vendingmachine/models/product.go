package models

type Product struct {
	name  string
	price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{name: name, price: price}
}
