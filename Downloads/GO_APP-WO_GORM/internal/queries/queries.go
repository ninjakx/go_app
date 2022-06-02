package queries

const (
	CreateDB = `
		CREATE TABLE IF NOT EXISTS users (
			id        int primary key,
			created_at timestamp,
			updated_at timestamp,
			deleted_at timestamp,
			username  text,
			password  text,
			first_name text,
			last_name  text,
			phone     text,
			status bool
		);

		CREATE TABLE IF NOT EXISTS user_addresses (
			id        int primary key,	   
			created_at timestamp,
			updated_at timestamp,
			deleted_at timestamp, 
			user_id       int,	
			address_line1 text, 
			address_line2 text,	
			city         text,	
			postal_code   text,	
			country      text,	
			phone        text,
			telephone    text,
			FOREIGN KEY (user_id) REFERENCES users(id)				
		);
	`
	QueryInsertUserData = `
		INSERT INTO users (id, created_at, updated_at, deleted_at, username, password, first_name, last_name, phone, status) VALUES(:id, :created_at, :updated_at, :deleted_at, :username, :password, :first_name, :last_name, :phone, :status)
	`
	QueryInsertAddrData = `
		INSERT INTO user_addresses (id, created_at, updated_at, deleted_at, user_id, address_line1, address_line2, city, postal_code, country, phone, telephone) VALUES(:id, :created_at, :updated_at, :deleted_at, :user_id, :address_line1, :address_line2, :city, :postal_code, :country, :phone, :telephone)
	`

	QueryFindUser = `SELECT * FROM users WHERE id=$1`

	QueryFindUserAddr = `SELECT * FROM user_addresses WHERE user_id=$1`

	QueryAlluser = `SELECT * FROM users`

	QueryFilterUserAddress = `SELECT * FROM user_addresses WHERE user_id=$1`
	
	QueryUpdateUser = `UPDATE users SET updated_at=:updated_at, username=:username, password=:password, first_name=:first_name, last_name=:last_name, phone=:phone, status=:status WHERE id=:id;`

	QueryFilterUserAddressWid = `SELECT * FROM user_addresses WHERE id=$2 AND user_id=$1;`

	QueryUpdateUserAddr = `UPDATE user_addresses SET updated_at=:updated_at, address_line1=:address_line1, address_line2=:address_line2, city=:city, postal_code=:postal_code, country=:country, phone=:phone, telephone=:telephone WHERE id=:id AND user_id=:user_id;`

	QueryUpdateUserStatus = `UPDATE users SET status=:status WHERE id=:id;`

	QueryDeleteUser = `DELETE FROM users WHERE id=$1;`

	QueryDeleteUserAddr = `DELETE FROM user_addresses WHERE id=:id AND user_id=:user_id;`

	QueryDeleteAddresses = `DELETE FROM user_addresses WHERE user_id=$1;`

)
