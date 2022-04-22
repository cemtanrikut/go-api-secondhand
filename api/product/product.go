package product

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Condition   string  `json:"condition"`
	Prize       float64 `json:"prize"`
	IsSold      bool    `json:"is_sold"`
	IsActive    bool    `json:"is_active"`
	CreateDate  string  `json:"create_date"`
}

func AddProduct(product Product) error {
	uuid := uuid.New()
	product.ProductID = uuid.String()
	//...
	return nil
}

func UpdateProduct(product Product) error {
	productID := product.ProductID
	updatedProduct := Product{
		ProductID:   productID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Condition:   product.Condition,
		Prize:       product.Prize,
		IsSold:      product.IsSold,
		IsActive:    product.IsActive,
	}
	fmt.Println(updatedProduct)
	return nil
}

func DeleteProduct(productID string) error {
	//is_active = false
	return nil
}

func GetProduct(productID string) (Product, error) {

	return Product{}, nil
}

func GetProductsByCustomer(customerID string) ([]Product, error) {

	return []Product{}, nil
}

func GetProductsByCategory(category string) ([]Product, error) {

	return []Product{}, nil
}

func GetProductsWithFilters(filter map[string]string) ([]Product, error) {

	return []Product{}, nil
}

func GetProducts() ([]Product, error) {

	return []Product{}, nil
}
