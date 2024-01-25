package model

type Product struct {
	_id         string
	name        string
	price       float32
	description string
	userID      string
	createDate  string
	updateDate  string
	isSold      bool
	isActive    bool
	isDeleted   bool
	isSpecial   bool
	onlineSale  bool
	category    string
	brand       string
	model       string
	quantity    int
}
