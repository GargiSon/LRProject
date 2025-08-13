package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"LRProject3/config"
	"LRProject3/models"
	"LRProject3/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/login.html")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	email := strings.ToLower(r.FormValue("email"))
	password := r.FormValue("password")

	if email == "" || password == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Email and password are required"})
		return
	}

	payload := map[string]string{"email": email, "password": password}
	jsonData, _ := json.Marshal(payload)

	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	loginURL := "https://api.loginradius.com/identity/v2/auth/login?apikey=" + apiKey

	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
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

	var lrResp models.LoginResponse
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &lrResp)

	sessionID := utils.RandomString(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", Value: "", Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func ValidateToken(token string) bool {
	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	url := "https://api.loginradius.com/identity/v2/auth/access_token/validate?apikey=" + apiKey + "&access_token=" + token

	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
