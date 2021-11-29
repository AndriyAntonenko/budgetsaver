package repository

import (
	"database/sql"
	"fmt"
)

type BudgetPostgres struct {
	db *sql.DB
}

func NewBudgetPostgres(db *sql.DB) *BudgetPostgres {
	return &BudgetPostgres{db}
}

func (r *BudgetPostgres) CreateBudget(payload CreateBudgetRecord) (*BudgetRecord, error) {
	var newRecord BudgetRecord
	createQuery := fmt.Sprintf("INSERT INTO %s (name, creator, finance_group_id, description) VALUES ($1, $2, $3, $4) RETURNING id, name, creator, finance_group_id, description;", budgetTable)
	row, err := r.db.Query(createQuery, payload.Name, payload.Creator, payload.FinanceGroupId, payload.Description)

	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.Scan(&newRecord.Id, &newRecord.Name, &newRecord.Creator, &newRecord.FinanceGroupId, &newRecord.Description)
	if err != nil {
		return nil, err
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}
