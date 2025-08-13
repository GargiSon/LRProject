package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"LRProject3/config"
	"LRProject3/models"
	"LRProject3/utils"
)

func apiHost() string {
	h := config.GetEnv("LOGINRADIUS_API_DOMAIN")
	if h == "" {
		h = "api.loginradius.com"
	}
	return h
}

func getSOTT() (string, error) {
	apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
	apiSecret := config.GetEnv("LOGINRADIUS_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		return "", fmt.Errorf("LOGINRADIUS_API_KEY or LOGINRADIUS_API_SECRET not set")
	}

	q := url.Values{}
	q.Set("apiKey", apiKey)
	q.Set("apiSecret", apiSecret)
	q.Set("timeDifference", "10") // minutes; keep small so SOTT doesn't expire

	sottURL := fmt.Sprintf("https://%s/identity/v2/manage/account/sott?%s", apiHost(), q.Encode())

	resp, err := http.Get(sottURL)
	if err != nil {
		return "", fmt.Errorf("request to SOTT API failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get SOTT (%d): %s", resp.StatusCode, string(body))
	}

	var data models.SottResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("decode SOTT response failed: %w", err)
	}
	if data.Sott == "" {
		return "", fmt.Errorf("SOTT missing in response")
	}
	return data.Sott, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetNoCacheHeaders(w)

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "static/register.html")
		return

	case http.MethodPost:
		var email, password, firstName, lastName string

		ct := r.Header.Get("Content-Type")
		if ct != "" && (ct == "application/json" || (len(ct) >= 16 && ct[:16] == "application/json")) {
			var body struct {
				Email     string `json:"email"`
				Password  string `json:"password"`
				FirstName string `json:"firstname"`
				LastName  string `json:"lastname"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
				email, password, firstName, lastName = body.Email, body.Password, body.FirstName, body.LastName
			}
		} else {
			_ = r.ParseForm()
			email = r.FormValue("email")
			password = r.FormValue("password")
			firstName = r.FormValue("firstname")
			lastName = r.FormValue("lastname")
		}

		apiKey := config.GetEnv("LOGINRADIUS_API_KEY")
		if apiKey == "" {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "LOGINRADIUS_API_KEY not set"})
			return
		}

		// 1) Get a fresh SOTT
		sott, err := getSOTT()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("SOTT error: %v", err)})
			return
		}

		// 2) Call register API
		lrURL := fmt.Sprintf("https://%s/identity/v2/auth/register?apiKey=%s", apiHost(), url.QueryEscape(apiKey))

		payload := map[string]any{
			"Email": []map[string]string{
				{"Type": "Primary", "Value": email},
			},
			"Password":  password,
			"FirstName": firstName,
			"LastName":  lastName,
		}
		reqBody, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", lrURL, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-LoginRadius-Sott", sott)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "Could not reach LoginRadius"})
			return
		}
		defer resp.Body.Close()

		respBytes, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(respBytes)
			return
		}

		var ok models.LrRegisterResponse
		_ = json.Unmarshal(respBytes, &ok)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message":   "Account created",
			"uid":       ok.Uid,
			"firstName": ok.FirstName,
			"lastName":  ok.LastName,
		})
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}
