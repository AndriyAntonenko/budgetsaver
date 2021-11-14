package dto

type UserSignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
