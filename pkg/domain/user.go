package domain

type User struct {
	Id string `db:"id" json:"id"`

	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `json:"password"`
}

type CreateUserRecord struct {
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Salt         string `db:"salt"`
}
