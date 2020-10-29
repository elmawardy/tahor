package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	// init DB connection & migrate required models

	var err error
	DB, err = gorm.Open(Config["driver"], Config["constring"])
	defer DB.Close()
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&User{})

	// end init DB

	r := mux.NewRouter()
	r.HandleFunc("/api/signin", Log(SignInHandler()))
	r.HandleFunc("/api/register", Log(RegisterHandler()))
	r.HandleFunc("/api/signout", Log(SignOutHandler()))
	r.HandleFunc("/api/getbasicinfo", Log(GetBasicInfoHandler()))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8089", r))
}
