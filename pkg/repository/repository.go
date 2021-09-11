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

type Group interface {
	CreateGroup(string, *domain.CreateGroupPayload) (*domain.Group, error)
}

type Repository struct {
	Authorization
	Group
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Group:         NewGroupPostgres(db),
	}
}
