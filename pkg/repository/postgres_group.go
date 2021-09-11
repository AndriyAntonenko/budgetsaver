package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

type GroupPostgres struct {
	db *sql.DB
}

func NewGroupPostgres(db *sql.DB) *GroupPostgres {
	return &GroupPostgres{db}
}

func (r *GroupPostgres) CreateGroup(userId string, payload *domain.CreateGroupPayload) (*domain.Group, error) {
	var result domain.Group

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create group
	query := fmt.Sprintf("INSERT INTO \"%s\" (name, type) VALUES ($1, $2) RETURNING id, name, type", groupTable)
	row := tx.QueryRowContext(ctx, query, payload.Name, payload.Type)
	err = row.Scan(&result.Id, &result.Name, &result.Type)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create relation
	query = fmt.Sprintf("INSERT INTO %s (user_id, group_id) VALUES ($1, $2)", usersGroupTable)
	_, err = tx.ExecContext(ctx, query, userId, result.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &result, nil
}
