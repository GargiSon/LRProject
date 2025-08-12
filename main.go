package main

import (
	"log"
	"net/http"

	"LRProject3/config"
	"LRProject3/db"
	"LRProject3/handlers"
)

func main() {
	config.LoadEnv()
	db.ConnectMongo()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/home", handlers.AuthMiddleware(handlers.HomeHandler))
	http.HandleFunc("/forgot", handlers.ForgotPasswordHandler)
	http.HandleFunc("/reset", handlers.ResetPasswordHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
