package repository

import (
	"database/sql"
	"fmt"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(payload domain.CreateUserRecord) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, salt) VALUES ($1, $2, $3, $4) RETURNING id;", usersTable)
	row, err := r.db.Query(query, payload.Name, payload.Email, payload.PasswordHash, payload.Salt)

	if err != nil {
		return "", err
	}

	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
