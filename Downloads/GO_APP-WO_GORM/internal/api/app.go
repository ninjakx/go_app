package api

import (
	"GO_APP/internal/handler"
	"GO_APP/internal/model"
	"GO_APP/config"
	"fmt"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// App has router and db instances
type App struct {
	Router *httprouter.Router//mux.Router
	DB     *sqlx.DB
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

	db, err := sqlx.Connect(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	} else {
		fmt.Printf("Connected to database\n")
	}

	a.DB = model.DBMigrate(db)
	a.Router = httprouter.New()
	a.setRouters()
}

// https://github.com/gin-gonic/gin/issues/1681
// Set all required routers
func (a *App) setRouters() {
	router := a.Router
	// Routing for handling the projects
	router.GET("/users", a.GetAllUser)
	router.GET("/users/:id", a.GetUser)
	router.GET("/users/:id/address", a.GetUserAddress)
	router.POST("/users", a.CreateUser)
	router.POST("/users/:id/add_address", a.CreateUserAddress)
	router.PUT("/users/:id/update_user", a.UpdateUser)
	router.PUT("/users/:id/update_address/:addr_id", a.UpdateUserAddress)
	router.PUT("/users/:id/disable", a.DisableUser)
	router.PUT("/users/:id/enable", a.EnableUser)
	router.DELETE("/users/:id", a.DeleteUser)
	router.DELETE("/users/:id/del/:addr_id", a.DeleteUserAddress)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.CreateUser(a.DB, w, r)
}

// Handlers to manage User Data
func (a *App) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.GetUser(a.DB, w, r, ps)
}

func (a *App) GetAllUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.GetAllUser(a.DB, w, r)
}

func (a *App) GetUserAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.GetUserAddress(a.DB, w, r, ps)
}

func (a *App) CreateUserAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.CreateUserAddress(a.DB, w, r, ps)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.UpdateUser(a.DB, w, r, ps)
}

func (a *App) UpdateUserAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.UpdateUserAddress(a.DB, w, r, ps)
}

func (a *App) DisableUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.DisableUser(a.DB, w, r, ps)
}

func (a *App) EnableUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.EnableUser(a.DB, w, r, ps)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.DeleteUser(a.DB, w, r, ps)
}

func (a *App) DeleteUserAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.DeleteUserAddress(a.DB, w, r, ps)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
