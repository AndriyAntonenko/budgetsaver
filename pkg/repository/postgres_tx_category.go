package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type TxCategoryPostgres struct {
	db *sql.DB
}

func NewTxCategoryPostgres(db *sql.DB) *TxCategoryPostgres {
	return &TxCategoryPostgres{db}
}

const columnsString = "id, name, creator, finance_group, created_at, deleted_at"

func (r *TxCategoryPostgres) AddTxCategory(payload CreateTxCategoryRecord) (*TxCategoryRecord, error) {
	var txCategory TxCategoryRecord

	insertColumnsString := "name"
	insertValuesString := "$1"
	valuesSlice := []interface{}{payload.Name}

	if payload.Creator != nil {
		insertColumnsString += ", creator"
		insertValuesString += ", $2"
		valuesSlice = append(valuesSlice, *payload.Creator)
	}

	if payload.FinanceGroup != nil {
		insertColumnsString += ", finance_group"
		insertValuesString += ", $3"
		valuesSlice = append(valuesSlice, *payload.FinanceGroup)
	}

	createQuery := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING %s;",
		budgetTxCategory,
		insertColumnsString,
		insertValuesString,
		columnsString,
	)
	row, err := r.db.Query(createQuery, valuesSlice...)

	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.Scan(&txCategory.Id, &txCategory.Name, &txCategory.Creator, &txCategory.FinanceGroup, &txCategory.CreatedAt, &txCategory.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &txCategory, nil
}

func (r *TxCategoryPostgres) GetTxCategoryById(categoryId string) (*TxCategoryRecord, error) {
	var txCategory TxCategoryRecord
	getQuery := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", columnsString, budgetTxCategory)

	row := r.db.QueryRow(getQuery, categoryId)
	err := row.Scan(&txCategory.Id, &txCategory.Name, &txCategory.Creator, &txCategory.FinanceGroup, &txCategory.CreatedAt, &txCategory.DeletedAt)
	if err != nil {
		return nil, err
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("budget with such id not found")
		}
		return nil, err
	}

	return &txCategory, nil
}
