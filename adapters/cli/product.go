package cli

import (
	"fmt"

	"github.com/mathsant/go-arch-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64, productQuantity int) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice, productQuantity)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name: %s has been created with price: %f, quantity: %d and status: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Quantity: %d\n Status: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus())
	}

	return result, nil
}
