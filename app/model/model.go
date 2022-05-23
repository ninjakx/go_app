package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	// Id	 		int  `sql:"type:int;primary key"`
	Username	string	`sql:"type:varchar(36)"`
	Password	string	`sql:"type:varchar(36)"`
	FirstName	string	`sql:"type:varchar(36)"`
	LastName 	string	`sql:"type:varchar(36)"`
	Phone		string	`sql:"type:varchar(36)"`
	// CreatedAt	string	`sql:"type:varchar(36)"`
	// ModifiedAt	string	`sql:"type:varchar(36)"`
	Status 			bool  	`sql:"type:bool;"`
	Addrs []UserAddress
}

type UserAddress struct {
	gorm.Model
	// Id	 			int    `sql:"type:int"`
	UserId			uint   `sql:"type:uint;"`
	AddressLine1	string `sql:"type:varchar(36)"`
	AddressLine2	string	`sql:"type:varchar(36)"`
	City		 	string	`sql:"type:varchar(36)"`
	PostalCode		string	`sql:"type:varchar(36)"`
	Country			string	`sql:"type:varchar(36)"`
	Phone			string	`sql:"type:varchar(36)"`
	Telephone		string	`sql:"type:varchar(36)"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{},&UserAddress{})
	return db
}

func (u *User) Disable() {
	u.Status = false
}

func (u *User) Enable() {
	u.Status = true
}