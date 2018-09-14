package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"github.com/gorilla/mux"
	"github.com/maxheckel/meditations/handlers"
	"net/http"
	"fmt"
	"github.com/golang-migrate/migrate/database/sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/Users/maxheckel/Go/src/github.com/maxheckel/meditations/db.sqlite")
	if err != nil{
		log.Fatal(err)
	}
	sqlite3.WithInstance(db, &sqlite3.Config{})
	fmt.Println("testing")
	r := mux.NewRouter()
	r.Handle("/", &handlers.ListPostsHandler{Database:db})
	http.ListenAndServe(":8080", r)
}
