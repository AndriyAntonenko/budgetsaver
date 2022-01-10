package repository

import (
	"database/sql"
	"fmt"
)

const budgetTxCreateColumns = "budget_id, title, description, \"from\", \"to\", amount, author, tx_time"
const budgetTXColumns = "id, budget_id, title, description, \"from\", \"to\", amount, author, tx_time, created_at, deleted_at, category"

type BudgetTXPostgres struct {
	db *sql.DB
}

func NewBudgetTXPostgres(db *sql.DB) *BudgetTXPostgres {
	return &BudgetTXPostgres{db}
}

func (r *BudgetTXPostgres) CreateBudgetTx(payload CreateBudgetTxRecord) (*BudgetTxRecord, error) {
	var newRecord BudgetTxRecord

	columnsString := budgetTxCreateColumns
	valuesString := "$1, $2, $3, $4, $5, $6, $7, $8"
	valuesSlice := []interface{}{payload.BudgetId, payload.Title, payload.Description, payload.From, payload.To, payload.Amount, payload.Author, payload.TxTime}

	if payload.Category != nil {
		columnsString += ", category"
		valuesString += ", $9"
		valuesSlice = append(valuesSlice, *payload.Category)
	}

	createQuery := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING %s;",
		budgetTxTable,
		columnsString,
		valuesString,
		budgetTXColumns,
	)
	row, err := r.db.Query(createQuery, valuesSlice...)

	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.Scan(&newRecord.Id, &newRecord.BudgetId, &newRecord.Title, &newRecord.Description, &newRecord.From, &newRecord.To, &newRecord.Amount, &newRecord.Author, &newRecord.TxTime, &newRecord.CreatedAt, &newRecord.DeletedAt, &newRecord.Category)
	if err != nil {
		return nil, err
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}
