package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func setNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/login.html")
		return
	}

	email := strings.ToLower(r.FormValue("email"))
	password := r.FormValue("password")

	payload := map[string]string{"email": email, "password": password}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to prepare request", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("LOGINRADIUS_API_KEY")
	if apiKey == "" {
		log.Fatal("LOGINRADIUS_API_KEY environment variable is not set")
	}

	loginURL := "https://api.loginradius.com/identity/v2/auth/login?apikey=" + apiKey

	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Login request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Login failed:", string(bodyBytes))
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	var lrResp LoginResponse
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &lrResp)

	sessionID := randomString(32)
	saveSession(sessionID, lrResp.AccessToken)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setNoCacheHeaders(w)
	http.ServeFile(w, r, "static/home.html")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie("session_id"); err == nil {
		deleteSession(c.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: "session_id", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		token := getSessionToken(c.Value)
		if token == "" || !validateToken(token) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func validateToken(token string) bool {
	apiKey := os.Getenv("LOGINRADIUS_API_KEY")
	url := "https://api.loginradius.com/identity/v2/auth/access_token/validate?apikey=" + apiKey + "&access_token=" + token

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
