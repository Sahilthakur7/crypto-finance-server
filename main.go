package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func handler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world")
}