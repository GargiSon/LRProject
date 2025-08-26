package models

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type LrRegisterResponse struct {
	Uid       string `json:"Uid"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type SottResponse struct {
	Sott           string `json:"Sott"`
	ExpirationTime string `json:"ExpirationTime"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}
