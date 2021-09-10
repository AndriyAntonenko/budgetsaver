package domain

type UserSignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRecord struct {
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Salt         string `db:"salt"`
}

type UserRecord struct {
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Salt         string `db:"salt"`
	Id           string `db:"id"`
}

type UserProfile struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
