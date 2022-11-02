package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {

	godotenv.Load()

	log.Println("Server will start at http://localhost:8080/")

	ConnectDatabase()

	route := mux.NewRouter()

	AddApproutes(route)

	log.Fatal(http.ListenAndServe(":8080", route))
	//http.HandleFunc("/", RenderHome)
	//http.HandleFunc("/login", RenderHome)
	//err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	
}
