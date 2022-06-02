package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	"GO_APP/internal/queries"

)

type User struct {
	ID        int	    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
	Username  string	`db:"username"`
	Password  string	`db:"password"`
	FirstName string	`db:"first_name"`
	LastName  string	`db:"last_name"`
	Phone     string	`db:"phone"`
	Status bool			`db:"status"`
	Addrs  []UserAddress
}

type UserAddress struct {
	ID        int	    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
	UserId       int 	`db:"user_id"`
	AddressLine1 string `db:"address_line1"`
	AddressLine2 string	`db:"address_line2"`
	City         string	`db:"city"`	
	PostalCode   string	`db:"postal_code"`
	Country      string	`db:"country"`
	Phone        string	`db:"phone"`
	Telephone    string	`db:"telephone"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *sqlx.DB) *sqlx.DB {	
	db.MustExec(queries.CreateDB)
	return db
}

func (u *User) Disable() {
	u.Status = false
}

func (u *User) Enable() {
	u.Status = true
}
