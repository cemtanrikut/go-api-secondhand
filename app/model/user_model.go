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
}
