package main

import (
	"fmt"
	"latihan_uts/controller"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/detailTransactions2", controller.GetDetailUserTransactions2).Methods("GET")
	router.HandleFunc("/v1/detailTransactionsUser/{id}", controller.GetDetailUserTransByID).Methods("GET")
	router.HandleFunc("/v1/deleteProduct/{id}", controller.DeleteSingleProduct).Methods("DELETE")
	router.HandleFunc("/v1/insertProduct", controller.InsertNewProducts).Methods("POST")
	router.HandleFunc("/v1/login", controller.Login).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to Port 8888")
	log.Println("Connected to Port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
