package models

type DtoSignupRequest struct {
	Username string `json:"Username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DtoLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
