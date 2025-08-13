package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"LRProject3/config"
)

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/forgot.html")
		return
	}

	email := strings.ToLower(r.FormValue("email"))
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	resetURL := config.GetEnv("RESET_PASSWORD_URL")
	emailTemplate := config.GetEnv("EMAIL_TEMPLATE")

	payload := map[string]string{"email": email}
	jsonData, _ := json.Marshal(payload)

	apiEndpoint := "https://api.loginradius.com/identity/v2/auth/password"
	apiReqURL := fmt.Sprintf("%s?apikey=%s&resetPasswordUrl=%s", apiEndpoint, apiKey, url.QueryEscape(resetURL))
	if emailTemplate != "" {
		apiReqURL += "&emailTemplate=" + url.QueryEscape(emailTemplate)
	}

	resp, err := http.Post(apiReqURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to send reset request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, "Error: "+string(body), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("If an account exists, a password link has been sent."))
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		token := r.URL.Query().Get("vtoken")
		if token == "" {
			token = r.URL.Query().Get("token")
		}
		if token == "" {
			http.Error(w, "Invalid or missing reset token", http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles("static/reset.html")
		if err != nil {
			http.Error(w, "Failed to load reset page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]string{"Token": token})
		return
	}

	token := r.FormValue("token")
	if token == "" {
		token = r.FormValue("vtoken")
	}
	newPassword := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if token == "" {
		http.Error(w, "Missing reset token", http.StatusBadRequest)
		return
	}
	if newPassword == "" || confirmPassword == "" {
		http.Error(w, "Password fields cannot be empty", http.StatusBadRequest)
		return
	}
	if newPassword != confirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	payload := map[string]string{
		"resetToken": token,
		"password":   newPassword,
	}
	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT",
		"https://api.loginradius.com/identity/v2/auth/password/reset?apikey="+apiKey,
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Password reset error: %s", string(body))
		http.Error(w, "Password reset failed. The link may be expired or invalid.", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
