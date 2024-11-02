package cli

import (
	"fmt"
	"github.com/rodrigospimentacwb/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productID string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f ans status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s  has been enabled.", res.GetName())

	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s  has been disabled.", res.GetName())
	default:
		res, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Status: %s\n", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
