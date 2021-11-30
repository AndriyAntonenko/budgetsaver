package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(payload CreateUserRecord) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, salt, last_login_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;", usersTable)
	row, err := r.db.Query(query, payload.Name, payload.Email, payload.PasswordHash, payload.Salt, payload.LastLoginAt)

	if err != nil {
		return "", err
	}

	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	err = row.Close()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (UserRecord, error) {
	return r.getUserRecordByField("email", email)
}

func (r *AuthPostgres) GetUserById(id string) (UserRecord, error) {
	return r.getUserRecordByField("id", id)
}

func (r *AuthPostgres) getUserRecordByField(field string, value string) (UserRecord, error) {
	var userRecord UserRecord
	query := fmt.Sprintf("SELECT id, name, email, password_hash, salt, last_login_at FROM %s WHERE %s = $1", usersTable, field)
	row := r.db.QueryRow(query, value)
	err := row.Scan(&userRecord.Id, &userRecord.Name, &userRecord.Email, &userRecord.PasswordHash, &userRecord.Salt, &userRecord.LastLoginAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return userRecord, errors.New("user with such email not found")
		}
		return userRecord, err
	}

	return userRecord, nil
}

func (r *AuthPostgres) UpdateLastLogin(userId string, lastLoginAt time.Time) (UserRecord, error) {
	var userRecord UserRecord
	query := fmt.Sprintf("UPDATE %s SET last_login_at = $1 WHERE id = $2 RETURNING id, name, email, password_hash, salt, last_login_at", usersTable)
	row := r.db.QueryRow(query, lastLoginAt, userId)
	err := row.Scan(&userRecord.Id, &userRecord.Name, &userRecord.Email, &userRecord.PasswordHash, &userRecord.Salt, &userRecord.LastLoginAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return userRecord, errors.New("user with such id not found")
		}
		return userRecord, err
	}

	return userRecord, nil
}
