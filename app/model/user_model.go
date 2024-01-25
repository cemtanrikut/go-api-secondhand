package model

type User struct {
	_id         int
	name        string
	surname     string
	password    string
	email       string
	phone       string
	dateofbirth string
	gender      bool
	createDate  string
	updateDate  string
	isDeleted   bool
	isActive    bool
	address     Address
}

type Address struct {
	userID    string
	country   string
	city      string
	town      string
	postCode  string
	lat       float64
	lon       float64
	isActive  bool
	isDeleted bool
	isPrimary bool
}
