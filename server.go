package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// RenderHome Rendering the Home Page
func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

func main() {

	//godotenv.Load()

	log.Println("Server will start at http://localhost:8080/")

	//ConnectDatabase()

	route := mux.NewRouter()

	log.Println("Loadeding Routes...")

	route.HandleFunc("/", RenderHome)
	route.HandleFunc("/login", RenderHome)

	route.HandleFunc("/login/facebook", InitFacebookLogin)

	route.HandleFunc("/facebook/callback", HandleFacebookLogin)

	route.HandleFunc("/userDetails", GetUserDetails).Methods("GET")

	log.Println("Routes are Loaded.")
	//AddApproutes(route)

	log.Fatal(http.ListenAndServe(":8080", route))
	//http.HandleFunc("/", RenderHome)
	//http.HandleFunc("/login", RenderHome)
	//err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)

}
