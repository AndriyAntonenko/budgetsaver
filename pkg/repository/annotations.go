package repository

import (
	"database/sql"
	"time"
)

// USER TABLE ANNOTATIONS
type CreateUserRecord struct {
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Salt         string `db:"salt"`
}

type UserRecord struct {
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	PasswordHash string       `db:"password_hash"`
	Salt         string       `db:"salt"`
	Id           string       `db:"id"`
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
