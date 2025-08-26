package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"LRProject3/config"
	"LRProject3/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	payload := map[string]string{
		"email":    loginReq.Email,
		"password": loginReq.Password,
	}
	jsonData, _ := json.Marshal(payload)

	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	loginURL := "https://devapi.lrinternal.com/identity/v2/auth/login?apikey=" + apiKey

	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Login request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	var lrResp models.LoginResponse
	if err := json.Unmarshal(bodyBytes, &lrResp); err != nil {
		http.Error(w, "Failed to parse login response", http.StatusInternalServerError)
		return
	}

	if lrResp.AccessToken == "" {
		http.Error(w, "Login failed: no access token", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "lr_token",
		Value:    lrResp.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("lr_token")
	if err == nil && cookie.Value != "" {
		apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
		logoutURL := "https://devapi.lrinternal.com/identity/v2/auth/access_token/invalidate?apikey=" + apiKey + "&access_token=" + cookie.Value

		client := &http.Client{}
		_, _ = client.Get(logoutURL)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "lr_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("lr_token")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if !ValidateToken(cookie.Value) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func ValidateToken(token string) bool {
	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	url := "https://devapi.lrinternal.com/identity/v2/auth/access_token/validate?apikey=" + apiKey + "&access_token=" + token

	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
