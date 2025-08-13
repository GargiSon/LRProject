package models

type LoginResponse struct {
	AccessToken string `json:"access_token"`
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
