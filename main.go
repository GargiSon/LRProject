package main

import (
	"log"
	"net/http"

	"LRProject3/config"
	"LRProject3/handlers"
)

func main() {
	config.LoadEnv()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/home", handlers.AuthMiddleware(handlers.HomeHandler))
	http.HandleFunc("/forgot", handlers.ForgotPasswordHandler)
	http.HandleFunc("/reset", handlers.ResetPasswordHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
