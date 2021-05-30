package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/models"
)

const (
	host    = "localhost"
	port    = 5432
	user    = "piero"
	db_name = "lenslocked_dev"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		db_name,
	)
	us, err := models.NewUserService(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer us.Close()
	// us.DestructiveReset()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	fmt.Println("Starting server on http://localhost:3000 ...")
	http.ListenAndServe(":3000", r)
}
