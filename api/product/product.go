package product

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
