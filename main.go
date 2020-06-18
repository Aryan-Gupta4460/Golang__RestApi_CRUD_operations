package main

import (
	"fmt"
	"net/http"

	"github.com/Aryan-Gupta4460/Product/apis"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/product/findall", apis.FindAll).Methods("GET")
	router.HandleFunc("/api/product/search/{keyword}", apis.Search).Methods("GET")
	router.HandleFunc("/api/product/search/{min}/{max}", apis.SearchPrices).Methods("GET")
	router.HandleFunc("/api/product/create", apis.Create).Methods("POST")
	router.HandleFunc("/api/product/update", apis.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}", apis.Delete).Methods("DELETE")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}

}
