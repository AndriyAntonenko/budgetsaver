package repository

import (
	"database/sql"
	"fmt"

	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	_ "github.com/lib/pq"
)

type tableName string

const (
	usersTable             tableName = "users"
	financeGroupTable      tableName = "finance_group"
	usersFinanceGroupTable tableName = "users_finance_group"
)

func NewPostgresDB(cnf config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.Username, cnf.Password, cnf.DBName, cnf.SSLMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
