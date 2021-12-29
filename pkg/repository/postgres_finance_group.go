package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
)

const allGroupColumns = "id, name, description, created_at, deleted_at"

type FinanceGroupPostgres struct {
	db *sql.DB
}

func NewFinanceGroupPostgres(db *sql.DB) *FinanceGroupPostgres {
	return &FinanceGroupPostgres{db}
}

func (r *FinanceGroupPostgres) CreateFinanceGroup(userId string, groupPayload CreateFinanceGroupRecord) (*FinanceGroupRecord, error) {
	var newRecord FinanceGroupRecord

	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// Creating a group
	createGroupQuery := fmt.Sprintf("INSERT INTO %s (name, description) VALUES ($1, $2) RETURNING %s;", financeGroupTable, allGroupColumns)
	groupRow, err := tx.QueryContext(ctx, createGroupQuery, groupPayload.Name, groupPayload.Description)
	if err != nil {
		return nil, err
	}

	groupRow.Next()
	// id, name, description, created_at, deleted_at
	err = groupRow.Scan(&newRecord.Id, &newRecord.Name, &newRecord.Description, &newRecord.CreatedAt, &newRecord.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = groupRow.Close()
	if err != nil {
		return nil, err
	}

	// assigning owner to group
	assingOwnerQuery := fmt.Sprintf("INSERT INTO %s (user_id, group_id, role) VALUES ($1, $2, $3);", usersFinanceGroupTable)
	_, err = tx.ExecContext(ctx, assingOwnerQuery, userId, newRecord.Id, domain.Owner)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func (r *FinanceGroupPostgres) GetUsersFinanceGroups(userId string) ([]FinanceGroupWithMemberCount, error) {
	groups := make([]FinanceGroupWithMemberCount, 0)
	query := fmt.Sprintf(`
		SELECT fg.*, COUNT(ufg.user_id) AS members_count
		FROM %s fg
		INNER JOIN %s ufg ON
			ufg.group_id = fg.id
			AND fg.id IN (
				SELECT
					ufg2.group_id
				FROM
					%s ufg2
				WHERE
					ufg2.user_id = $1
			)
		GROUP BY
			fg.id,
			fg.name ,
			fg.description ,
			fg.created_at ,
			fg.deleted_at ;
	`, financeGroupTable, usersFinanceGroupTable, usersFinanceGroupTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		financeGroup := new(FinanceGroupWithMemberCount)
		err = rows.Scan(&financeGroup.Id, &financeGroup.Name, &financeGroup.Description, &financeGroup.CreatedAt, &financeGroup.DeletedAt, &financeGroup.MemberCount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, *financeGroup)
	}

	return groups, nil
}

func (r *FinanceGroupPostgres) GetUserRoleInFinanceGroup(groupId string, userId string) (domain.FinanceGroupRole, error) {
	var role domain.FinanceGroupRole
	dbQuery := fmt.Sprintf("SELECT \"role\" FROM %s WHERE group_id = $1 AND user_id = $2", usersFinanceGroupTable)
	row := r.db.QueryRow(dbQuery, groupId, userId)
	err := row.Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return role, nil
}
