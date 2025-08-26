package handlers

import (
	"LRProject3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lr_token")
	if err != nil {
		http.Error(w, "Unauthorized - no token found", http.StatusUnauthorized)
		return
	}
	accessToken := cookie.Value

	apiKey := os.Getenv("LOGINRADIUS_API_KEY")
	if apiKey == "" {
		http.Error(w, "Server misconfigured: missing API key", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf(
		"https://devapi.lrinternal.com/identity/v2/auth/account?apikey=%s&access_token=%s",
		apiKey,
		accessToken,
	)

	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error contacting LoginRadius", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err != nil {
		http.Error(w, "Error parsing profile", http.StatusInternalServerError)
		return
	}

	profile := models.ProfileResponse{}
	if v, ok := raw["FirstName"].(string); ok {
		profile.FirstName = v
	}
	if v, ok := raw["LastName"].(string); ok {
		profile.LastName = v
	}
	if emails, ok := raw["Email"].([]any); ok && len(emails) > 0 {
		if emailObj, ok := emails[0].(map[string]any); ok {
			if val, ok := emailObj["Value"].(string); ok {
				profile.Email = val
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
