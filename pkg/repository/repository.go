package repository

import (
	"database/sql"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
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
	// pass finance group id and user id
	GetUserRoleInFinanceGroup(string, string) (domain.FinanceGroupRole, error)
}

type Budget interface {
	CreateBudget(CreateBudgetRecord) (*BudgetRecord, error)
}

type Repository struct {
	Authorization
	FinanceGroup
	Budget
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		FinanceGroup:  NewFinanceGroupPostgres(db),
		Budget:        NewBudgetPostgres(db),
	}
}
