package main

import (
	"log"
	"net/http"

	"github.com/MMMarinov/GoProjects/tree/Project3-BookstoreManagementAPI/Project3-BookManagementAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
