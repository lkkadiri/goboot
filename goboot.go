package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lkkadiri/goboot/config"
	"github.com/lkkadiri/goboot/controllers"
	"github.com/lkkadiri/goboot/sockets"
)

// HomeHandler router
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GoBoot Home\n"))
}

func main() {

	fmt.Printf("Starting GoBoot Server \n")
	fmt.Printf("Server is running at http://localhost:8000\n")

	//Get DB
	db := dbconfig.Init()

	// Mux and Route Handling
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	// Get a UserController instance
	uc := controllers.NewUserController(db)
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("", uc.GetAllUsers).Methods("GET")
	s.HandleFunc("", uc.CreateUser).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", uc.GetUser).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", uc.RemoveUser).Methods("DELETE")

	ws := r.PathPrefix("/users/ws").Subrouter()
	ws.HandleFunc("", socket.WsHandler).Methods("GET")
	// Bind to a port and pass our router in
	log.Print(http.ListenAndServe(":8000", r))

}
