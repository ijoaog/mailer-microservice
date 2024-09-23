package main

import (
	"log"
	"mailer-microservice/config" // Adicione o pacote de configuração

	"mailer-microservice/controllers"
	"mailer-microservice/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	config.Init()

	r.Handle("/send-mail", middleware.AuthMiddleware(http.HandlerFunc(controllers.SendMailHandler))).Methods("POST")

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
