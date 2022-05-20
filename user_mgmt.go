package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	// "time"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	// gorm.Model
	Id	 		int  `sql:"type:int;primary key"`
	Username	string	`sql:"type:varchar(36)"`
	Password	string	`sql:"type:varchar(36)"`
	FirstName	string	`sql:"type:varchar(36)"`
	LastName 	string	`sql:"type:varchar(36)"`
	Phone		string	`sql:"type:varchar(36)"`
	CreatedAt	string	`sql:"type:varchar(36)"`
	ModifiedAt	string	`sql:"type:varchar(36)"`
	Addrs []UserAddress
}

type UserAddress struct {
	// gorm.Model
	Id	 			int `sql:"type:int"`
	UserId			int `sql:"type:int;primary key"`
	AddressLine1	string `sql:"type:varchar(36)"`
	AddressLine2	string	`sql:"type:varchar(36)"`
	City		 	string	`sql:"type:varchar(36)"`
	PostalCode		string	`sql:"type:varchar(36)"`
	Country			string	`sql:"type:varchar(36)"`
	Phone			string	`sql:"type:varchar(36)"`
	Telephone		string	`sql:"type:varchar(36)"`
}

// type UserPayment struct {
// 	gorm.Model
// 	Id	 			int
// 	UserId			int
// 	PaymentType		string
// 	Provider		string
// 	AccountNo	 	int
// 	Expiry			string
// }

var (
	users = []User{
		{Id: 1, Username: "asd", Password: "ghghgh", FirstName: "aa", LastName: "bb",
			 Phone: "123", CreatedAt: "2006-01-02 15:04:05", ModifiedAt: "2006-01-02 15:04:05"},
		{Id: 2, Username: "B1", Password: "bc", FirstName: "bb", LastName: "cc",
			 Phone: "1234", CreatedAt: "2007-01-02 15:04:05", ModifiedAt: "2007-01-02 15:04:05"},
		{Id: 3, Username: "C1", Password: "cc", FirstName: "cc", LastName: "dd",
			 Phone: "1235", CreatedAt: "2008-01-02 15:04:05", ModifiedAt: "2008-01-02 15:04:05"},
	}


	addresses = []UserAddress{
		{Id: 1, UserId: 1, AddressLine1: "a123", AddressLine2: "b123",
			 City: "luck", PostalCode: "2290", Country: "India", Phone: "345", Telephone: "678"},
		{Id: 2, UserId: 2, AddressLine1: "h123", AddressLine2: "hasa123",
			 City: "lucknow", PostalCode: "2290", Country: "India", Phone: "345", Telephone: "678"},
		{Id: 3, UserId: 3, AddressLine1: "dsds123", AddressLine2: "bdsd123",
			 City: "KNPR", PostalCode: "229340", Country: "India", Phone: "323245", Telephone: "678"},
	}


	// payments = []UserPayment{
	// 	{Id: 1, UserId: 1, PaymentType: "a123", Provider: "b123",
	// 		 AccountNo: 1456, Expiry: "2290"},
	// 	{Id: 2, UserId: 2, PaymentType: "h123", Provider: "hasa123",
	// 		 AccountNo: 123, Expiry: "2290"},
	// 	{Id: 3, UserId: 3, PaymentType: "dsds123", Provider: "bdsd123",
	// 		 AccountNo: 789, Expiry: "229340"},
	// }
	// users = []User{
	// 	{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
	// 	{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
	// 	{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
	// 	{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	// }
)

type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

var db *gorm.DB
var err error

var (
	drivers = []Driver{
		{Name: "Jimmy Johnson", License: "ABC123"},
		{Name: "Howard Hills", License: "XYZ789"},
		{Name: "Craig Colbin", License: "DEF333"},
	}
	cars = []Car{
		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)

// config postgressql DB
const (
    host     = "localhost"
    port     = 5432
    user     = "kriti"
    password = "nkx01"
    dbname   = "go_dummy"
)


func main() {
	router := mux.NewRouter()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlconn)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&User{}, &UserAddress{})


	// db.AutoMigrate(&Driver{})
	// db.AutoMigrate(&Car{})



	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&UserAddress{})
	// db.AutoMigrate(&UserPayment{})


	// for index := range cars {
	// 	db.Create(&cars[index])
	// }

	// for index := range drivers {
	// 	db.Create(&drivers[index])
	// }

	for index := range users {
		// fmt.Println(index, "==", &users[index])
		fmt.Println(db.NewRecord(users[index].Id))
		db.FirstOrCreate(&users[index])
		// if db.NewRecord(users[index].Id)==true{
		// 	db.Create(&users[index])
		// }
	}

	for index := range addresses {
		// if db.NewRecord(addresses[index].Id)==true{
			// db.Create(&addresses[index])
		// }
		db.FirstOrCreate(&addresses[index])

	}

	// for index := range payments {
	// 	db.Create(&payments[index])
	// }


	// router.HandleFunc("/cars", GetCars).Methods("GET")
	// router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	// router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	// router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("#########")
	// var driver Driver
	// var cars []Car
	// db.First(&driver, params["id"])
	// db.Model(&driver).Related(&cars)
	// driver.Cars = cars
	// json.NewEncoder(w).Encode(&driver)
	params := mux.Vars(r)
	var user User
	// var addresses []UserAddress // User->users

	err := db.First(&user, "id = ?", params["id"]).Error
	if err == nil {
	json1, _ := json.Marshal(user)
	fmt.Println(string(json1))
	} else {
		fmt.Printf("!!!!!")
	}
	

	// result := db.First(&user, "Id = ?", params["id"])
	// fmt.Println(result)
	// db.Model(&user).Related(&addresses)
	// user.Addrs = addresses
	json.NewEncoder(w).Encode(&user)





	// db.Model(&addresses{}).Select("UserAddress.UserId").Joins("left join User on User.Id = UserAddress.UserId").Scan(&user{})
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	
	// json.NewEncoder(w).Encode(&user)

	// db.First(&user, params["id"])
	// db.Model(&user).carsRelated(&cars)
	// driver.Cars = cars
	// db.Where("UserId = ?", user.UserId).Model(&Role{}).Find(&roles)
	// db.Joins("join UserAddress ua ON ua.UserId = User.Id").Find(&users)
	// json.NewEncoder(w).Encode(&driver)
}

// func GetCars(w http.ResponseWriter, r *http.Request) {
// 	var cars []Car
// 	db.Find(&cars)
// 	json.NewEncoder(w).Encode(&cars)
// }

// func GetCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var car Car
// 	db.First(&car, params["id"])
// 	json.NewEncoder(w).Encode(&car)
// }

// func GetDriver(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var driver Driver
// 	var cars []Car
// 	db.First(&driver, params["id"])
// 	db.Model(&driver).Related(&cars)
// 	driver.Cars = cars
// 	json.NewEncoder(w).Encode(&driver)
// }

// func DeleteCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var car Car
// 	db.First(&car, params["id"])
// 	db.Delete(&car)

// 	var cars []Car
// 	db.Find(&cars)
// 	json.NewEncoder(w).Encode(&cars)
// }



// type Driver struct {
// 	gorm.Model
// 	Name    string
// 	License string
// 	Users    []User
// }

// type User struct {
// 	gorm.Model
// 	Year      int
// 	Make      string
// 	ModelName string
// 	DriverID  int
// }

// var db *gorm.DB
// var err error

// var (
// 	drivers = []Driver{
// 		{Name: "Jimmy Johnson", License: "ABC123"},
// 		{Name: "Howard Hills", License: "XYZ789"},
// 		{Name: "Craig Colbin", License: "DEF333"},
// 	}
// 	users = []User{
// 		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
// 		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
// 		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
// 		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
// 	}
// )

// // config postgressql DB
// const (
//     host     = "localhost"
//     port     = 5432
//     user     = "kriti"
//     password = "nkx01"
//     dbname   = "go_dummy"
// )


// func main() {
// 	router := mux.NewRouter()

// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = gorm.Open("postgres", psqlconn)

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	defer db.Close()

// 	db.AutoMigrate(&Driver{})
// 	db.AutoMigrate(&User{})

// 	for index := range users {
// 		db.Create(&users[index])
// 	}

// 	for index := range drivers {
// 		db.Create(&drivers[index])
// 	}

// 	router.HandleFunc("/users", GetCars).Methods("GET")
// 	router.HandleFunc("/users/{id}", GetCar).Methods("GET")
// 	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
// 	router.HandleFunc("/users/{id}", DeleteCar).Methods("DELETE")

// 	handler := cors.Default().Handler(router)

// 	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
// }

// func GetCars(w http.ResponseWriter, r *http.Request) {
// 	var users []User
// 	db.Find(&users)
// 	json.NewEncoder(w).Encode(&users)
// }

// func GetCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var user User
// 	db.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(&user)
// }

// func GetDriver(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var driver Driver
// 	var users []User
// 	db.First(&driver, params["id"])
// 	db.Model(&driver).Related(&users)
// 	driver.Users = users
// 	json.NewEncoder(w).Encode(&driver)
// }

// func DeleteCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var user User
// 	db.First(&user, params["id"])
// 	db.Delete(&user)

// 	var users []User
// 	db.Find(&users)
// 	json.NewEncoder(w).Encode(&users)
// }