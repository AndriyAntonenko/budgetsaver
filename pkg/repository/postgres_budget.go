package repository

import (
	"database/sql"
	"errors"
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
	createQuery := fmt.Sprintf("INSERT INTO %s (name, creator, finance_group_id, description, currency) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, creator, finance_group_id, description, currency;", budgetTable)
	row, err := r.db.Query(createQuery, payload.Name, payload.Creator, payload.FinanceGroupId, payload.Description, payload.Currency)

	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.Scan(&newRecord.Id, &newRecord.Name, &newRecord.Creator, &newRecord.FinanceGroupId, &newRecord.Description, &newRecord.Currency)
	if err != nil {
		return nil, err
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func (r *BudgetPostgres) GetUserBudget(budgetId string, userId string) (*BudgetRecord, error) {
	var budgetRecord BudgetRecord
	query := fmt.Sprintf(`
		SELECT
			b.id,
			b.name,
			b.creator,
			b.finance_group_id,
			b.description,
			b.currency
		FROM %s b
		INNER JOIN %s fg ON fg.id = b.finance_group_id AND b.id = $1
		INNER JOIN %s ufg ON ufg.group_id = fg.id AND ufg.user_id = $2;
	`, budgetTable, financeGroupTable, usersFinanceGroupTable)

	row := r.db.QueryRow(query, budgetId, userId)
	err := row.Scan(&budgetRecord.Id, &budgetRecord.Name, &budgetRecord.Creator, &budgetRecord.FinanceGroupId, &budgetRecord.Description, &budgetRecord.Currency)

	if err != nil {
		if err == sql.ErrNoRows {
			return &budgetRecord, errors.New("budget with such id not found")
		}
		return &budgetRecord, err
	}

	return &budgetRecord, nil
}
