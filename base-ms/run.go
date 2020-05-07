package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/endpoint").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called this thing!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
