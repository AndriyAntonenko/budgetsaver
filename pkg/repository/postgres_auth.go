package repository

import (
	"database/sql"
	"errors"
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

func (r *AuthPostgres) GetUserByEmail(email string) (domain.UserRecord, error) {
	return r.getUserRecordByField("email", email)
}

func (r *AuthPostgres) GetUserById(id string) (domain.UserRecord, error) {
	return r.getUserRecordByField("id", id)
}

func (r *AuthPostgres) getUserRecordByField(field string, value string) (domain.UserRecord, error) {
	var userRecord domain.UserRecord
	query := fmt.Sprintf("SELECT id, name, email, password_hash, salt FROM %s WHERE %s = $1", usersTable, field)
	row := r.db.QueryRow(query, value)
	err := row.Scan(&userRecord.Id, &userRecord.Name, &userRecord.Email, &userRecord.PasswordHash, &userRecord.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return userRecord, errors.New("user with such email not found")
		}
		return userRecord, err
	}

	return userRecord, nil
}
