package dto

import "github.com/mathsant/go-arch-hexagonal/application"

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
	Quantity int     `json:"quantity"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	product.Quantity = p.Quantity

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
