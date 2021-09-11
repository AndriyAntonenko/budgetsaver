package repository

import (
	"database/sql"
	"fmt"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

type BudgetPostgres struct {
	db *sql.DB
}

func NewBudgetPostgres(db *sql.DB) *BudgetPostgres {
	return &BudgetPostgres{db}
}

func (r *BudgetPostgres) CreateBudget(userId string, payload *domain.CreateBudgetPayload) (*domain.Budget, error) {
	var result domain.Budget
	query := fmt.Sprintf("INSERT INTO %s (name, \"description\", group_id) VALUES ($1, $2, $3) RETURNING id, name, \"description\";", budgetTable)
	row := r.db.QueryRow(query, payload.Name, payload.Description, payload.GroupId)
	err := row.Scan(&result.Id, &result.Name, &result.Description)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
