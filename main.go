package main

import (
	"api-go-deploy-heroku/app"
	"api-go-deploy-heroku/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

}
