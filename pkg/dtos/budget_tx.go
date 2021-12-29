package dto

import "time"

type CreateBudgetTxDto struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
}

type BudgetTxDto struct {
	Id          string    `db:"id"`
	BudgetId    string    `json:"budgetId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Amount      float64   `json:"amount"`
	Author      string    `json:"author"`
	TxTime      time.Time `json:"txTime"`
}
