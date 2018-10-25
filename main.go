package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go_simurgh/views"
	)

func main() {

	var addrport string = "localhost:8080"
	router := mux.NewRouter()
	Dispatcher(router)
	log.Printf("Starting server at %s\nQuit the server with CONTROL-C\n", addrport)
	log.Fatal(http.ListenAndServe(addrport, router))

}

func Dispatcher(router *mux.Router){
	router.HandleFunc("/", views.Index).Methods("GET")
	router.HandleFunc("/SignUp", views.Signup).Methods("POST")
}
