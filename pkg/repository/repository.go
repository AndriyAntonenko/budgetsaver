package repository

import (
	"database/sql"
)

type Authorization interface {
	CreateUser(CreateUserRecord) (string, error)
	GetUserByEmail(string) (UserRecord, error)
	GetUserById(string) (UserRecord, error)
}

type FinanceGroup interface {
	// pass user id and payload
	CreateFinanceGroup(string, CreateFinanceGroupRecord) (*FinanceGroupRecord, error)
	// pass user id
	GetUsersFinanceGroups(string) ([]FinanceGroupWithMemberCount, error)
}

type Repository struct {
	Authorization
	FinanceGroup
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		FinanceGroup:  NewFinanceGroupPostgres(db),
	}
}
