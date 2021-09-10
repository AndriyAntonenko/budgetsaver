package repository

import (
	"database/sql"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

type Authorization interface {
	CreateUser(domain.CreateUserRecord) (string, error)
	GetUserByEmail(string) (domain.UserRecord, error)
	GetUserById(string) (domain.UserRecord, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
