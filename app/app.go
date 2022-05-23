package app

import (
	"fmt"
	"log"
	"net/http"
	"GO_APP/app/handler"
	"GO_APP/app/model"
	"GO_APP/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.DBname,	
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	} else {
		fmt.Printf("Connected to database\n")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}


// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/users", a.GetAllUser)
	a.Get("/users/{id}", a.GetUser)
	a.Get("/users/{id}/address", a.GetUserAddress)
	a.Post("/users", a.CreateUser)
	a.Post("/users/{id}/add_address", a.CreateUserAddress)
	a.Put("/users/{id}/update_user", a.UpdateUser)
	a.Put("/users/{id}/{addr_id}/update_address", a.UpdateUserAddress)
	a.Put("/users/{id}/disable", a.DisableUser)
	a.Put("/users/{id}/enable", a.EnableUser)
	a.Delete("/users/{id}", a.DeleteUser)
	a.Delete("/users/{id}/{addr_id}", a.DeleteUserAddress)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}


// Handlers to manage Employee Data
func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}


func (a *App) GetAllUser(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUser(a.DB, w, r)
}


func (a *App) GetUserAddress(w http.ResponseWriter, r *http.Request) {
	handler.GetUserAddress(a.DB, w, r)
}


func (a *App) CreateUserAddress(w http.ResponseWriter, r *http.Request) {
	handler.CreateUserAddress(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUserAddress(a.DB, w, r)
}

func (a *App) DisableUser(w http.ResponseWriter, r *http.Request) {
	handler.DisableUser(a.DB, w, r)
}

func (a *App) EnableUser(w http.ResponseWriter, r *http.Request) {
	handler.EnableUser(a.DB, w, r)
}


func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}


func (a *App) DeleteUserAddress(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUserAddress(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}