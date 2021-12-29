package repository

import (
	"database/sql"
	"time"
)

// USER TABLE ANNOTATIONS
type CreateUserRecord struct {
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Salt         string    `db:"salt"`
	LastLoginAt  time.Time `db:"last_login_at"`
}

type UserRecord struct {
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	PasswordHash string       `db:"password_hash"`
	Salt         string       `db:"salt"`
	Id           string       `db:"id"`
	LastLoginAt  sql.NullTime `db:"last_login_at"`
	CreatedAt    time.Time    `db:"created_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

// GROUP TABLE ANNOTATIONS

type CreateFinanceGroupRecord struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}

type FinanceGroupRecord struct {
	Id          string       `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type FinanceGroupWithMemberCount struct {
	Id          string       `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	MemberCount int          `db:"members_count"`
}

// BUDGET TABLE ANNOTATIONS

type CreateBudgetRecord struct {
	Creator        string `db:"creator"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	FinanceGroupId string `db:"finance_group_id"`
}

type BudgetRecord struct {
	Id             string `db:"id"`
	Creator        string `db:"creator"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	FinanceGroupId string `db:"finance_group_id"`
}

// BUDGET TX TABLE ANNOTATIONS

type BudgetTxRecord struct {
	Id          string       `db:"id"`
	BudgetId    string       `db:"budget_id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	From        string       `db:"from"`
	To          string       `db:"to"`
	Amount      float64      `db:"amount"`
	Author      string       `db:"author"`
	TxTime      time.Time    `db:"tx_time"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type CreateBudgetTxRecord struct {
	BudgetId    string    `db:"budget_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	From        string    `db:"from"`
	To          string    `db:"to"`
	Amount      float64   `db:"amount"`
	Author      string    `db:"author"`
	TxTime      time.Time `db:"tx_time"`
}
