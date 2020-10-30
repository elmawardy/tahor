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
	DB.AutoMigrate(&Case{}, &CaseTag{})

	// end init DB

	r := mux.NewRouter()
	r.HandleFunc("/api/add", Log(AddCaseHandler()))
	r.HandleFunc("/api/get", Log(GetCasesHandler()))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8089", r))
}
