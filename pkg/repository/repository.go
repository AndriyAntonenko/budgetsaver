package repository

import (
	"database/sql"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

type Authorization interface {
	CreateUser(domain.CreateUserRecord) (string, error)
}

// TODO: Implement repository
type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
