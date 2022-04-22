package customer

import (
	"fmt"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerID   string   `json:"customer_id"`
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Phone        string   `json:"phone"`
	CustomerType bool     `json:"customer_type"` //Person or shop
	Location     Location `json:"location"`
	IsActive     bool     `json:"is_active"`
}

type Location struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func AddCustomer(customer Customer) error {
	uuid := uuid.New()
	customer.CustomerID = uuid.String()
	return nil
}

func UpdateCustomer(customer Customer) error {
	customerID := customer.CustomerID
	updatedCustomer := Customer{
		CustomerID:   customerID,
		Name:         customer.Name,
		Surname:      customer.Surname,
		Email:        customer.Email,
		Password:     customer.Password,
		Phone:        customer.Phone,
		CustomerType: customer.CustomerType,
		Location:     customer.Location,
	}
	fmt.Println(updatedCustomer)
	return nil
}

func DeleteCustomer(customerID string) error {
	//is_active = false
	return nil
}

func GetCustomer(customerID string) (Customer, error) {

	return Customer{}, nil
}
