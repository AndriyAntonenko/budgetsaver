package repository

import (
	"database/sql"
	"fmt"
)

const budgetTxCreateColumns = "budget_id, title, description, \"from\", \"to\", amount, author, tx_time"
const budgetTXColumns = "id, budget_id, title, description, \"from\", \"to\", amount, author, tx_time, created_at, deleted_at"

type BudgetTXPostgres struct {
	db *sql.DB
}

func NewBudgetTXPostgres(db *sql.DB) *BudgetTXPostgres {
	return &BudgetTXPostgres{db}
}

func (r *BudgetTXPostgres) CreateBudgetTx(payload CreateBudgetTxRecord) (*BudgetTxRecord, error) {
	var newRecord BudgetTxRecord
	createQuery := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING %s;",
		budgetTxTable,
		budgetTxCreateColumns,
		budgetTXColumns,
	)
	row, err := r.db.Query(createQuery, payload.BudgetId, payload.Title, payload.Description, payload.From, payload.To, payload.Amount, payload.Author, payload.TxTime)

	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.Scan(&newRecord.Id, &newRecord.BudgetId, &newRecord.Title, &newRecord.Description, &newRecord.From, &newRecord.To, &newRecord.Amount, &newRecord.Author, &newRecord.TxTime, &newRecord.CreatedAt, &newRecord.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}
