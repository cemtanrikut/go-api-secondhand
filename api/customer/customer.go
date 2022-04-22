package customer

type Customer struct {
	CustomerID   string   `json:"customer_id"`
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Phone        string   `json:"phone"`
	CustomerType bool     `json:"customer_type"` //Person or shop
	Location     Location `json:"location"`
}

type Location struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}
